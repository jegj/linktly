package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Env    *EnvVar
	router *chi.Mux
	Store  *Store
	// Slog var to allow change level on-the-fly
	ServerLogVar *slog.LevelVar
}

func NewServer(ctx context.Context) *Server {
	env := NewEnvVar()
	store, err := NewStore(ctx, env.GetDBConnectionString())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	server := &Server{
		router:       chi.NewRouter(),
		Env:          env,
		ServerLogVar: SetUpLogger(env),
		Store:        store,
	}
	server.routes()
	return server
}

func (s *Server) Start(ctx context.Context, serverStopCtx context.CancelFunc) {
	server := http.Server{
		Addr:         s.Env.ServerAddress,
		Handler:      s.router,
		IdleTimeout:  s.Env.IdleTimeout,
		ReadTimeout:  s.Env.ReadTimeout,
		WriteTimeout: s.Env.WriteTimeout,
	}

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, shutdownCancelFunc := context.WithTimeout(ctx, s.Env.ShutdownGracePeriod)
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
