package bootstrap

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
)

type Server struct {
	Env *EnvVar
	// Mongo mongo.Client
	router *chi.Mux
}

func NewServer() *Server {
	env := NewEnvVar()
	server := &Server{
		router: chi.NewRouter(),
		Env:    env,
	}

	server.routes()

	// app.Mongo = NewMongoDatabase(app.Env)
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
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		log.Println("Shutting down the server gracefully...")
		err := server.Shutdown(shutdownCtx)
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
