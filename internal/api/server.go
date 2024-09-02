package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/config"
	"github.com/jegj/linktly/internal/store"
)

type Server struct {
	config config.Config
	router *chi.Mux
	store  *store.Store
	// Slog var to allow change level on-the-fly
	serverLogVar *slog.LevelVar
}

func NewServer(cfg config.Config, ctx context.Context) *Server {
	store, err := store.NewStore(ctx, cfg.GetDBConnectionString())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	server := &Server{
		router:       chi.NewRouter(),
		config:       cfg,
		serverLogVar: SetUpLogger(cfg),
		store:        store,
	}
	server.routes()
	return server
}

func (s *Server) Start(ctx context.Context, serverStopCtx context.CancelFunc) {
	server := http.Server{
		Addr:         s.config.ServerAddress,
		Handler:      s.router,
		IdleTimeout:  s.config.IdleTimeout,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
	}

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, shutdownCancelFunc := context.WithTimeout(ctx, s.config.ShutdownGracePeriod)
		defer shutdownCancelFunc()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				slog.Error("graceful shutdown timed out.. forcing exit.")
				os.Exit(1)
			}
		}()

		// Trigger graceful shutdown
		slog.Info("Shutting down the server gracefully...")
		// Closing the db connection
		s.store.Close()
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		serverStopCtx()
	}()

	// Run the server
	slog.Info(fmt.Sprintf("Starting server on %v", os.Getenv("SERVER_ADDRESS")))
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Wait for server context to be stopped
	<-ctx.Done()
}
