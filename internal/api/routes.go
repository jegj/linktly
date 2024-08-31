package api

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jegj/linktly/internal/api/domain/accounts"
)

func (s *Server) routes() {
	if s.config.LogHttpRequest {
		s.router.Use(middleware.Logger)
	}
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.CleanPath)
	s.router.Use(middleware.Heartbeat("/ping"))

	// TODO: Pass dbstore
	accounts.LoadRoutes(s.router, s.config)
	// maintance.LoadRoutes(s.router, s.ServerLogVar)
}
