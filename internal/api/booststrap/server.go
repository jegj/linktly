package bootstrap

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	// Env *Env
	// Mongo mongo.Client
	router *chi.Mux
}

func NewServer() *Server {
	server := &Server{
		router: chi.NewRouter(),
	}

	server.routes()

	// app.Env = NewEnv()
	// app.Mongo = NewMongoDatabase(app.Env)
	return server
}

func (s *Server) Start(ctx context.Context, serverStopCtx context.CancelFunc) {
	server := http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: s.router,
		/*
			IdleTimeout:  s.cfg.IdleTimeout,
			ReadTimeout:  s.cfg.ReadTimeout,
			WriteTimeout: s.cfg.WriteTimeout,
		*/
	}

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownGracePeriod, err := strconv.Atoi(os.Getenv("SHUTDOWN_GRACE_PERIOD_SECONDS"))
		if err != nil {
			panic(errors.New("Invalid value for SHUTDOWN_GRACE_PERIOD"))
		}
		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, shutdownCancelFunc := context.WithTimeout(ctx, time.Duration(shutdownGracePeriod)*time.Second)
		defer shutdownCancelFunc()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err = server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	log.Printf("Starting server on %v", os.Getenv("SERVER_ADDRESS"))
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-ctx.Done()
}
