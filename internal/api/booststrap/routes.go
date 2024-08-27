package bootstrap

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jegj/linktly/internal/api/domain/accounts"
)

func (s *Server) routes() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Heartbeat("/ping"))

	s.router.Mount("/api", accounts.Routes())

	/*
		s.router.Route("/api/movies", func(r chi.Router) {
			r.Get("/", s.handleListMovies)
			r.Post("/", s.handleCreateMovie)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", s.handleGetMovie)
				r.Put("/", s.handleUpdateMovie)
				r.Delete("/", s.handleDeleteMovie)
			})
		})
	*/
}
