package bootstrap

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jegj/linktly/internal/api/domain/accounts"
	"github.com/jegj/linktly/internal/api/domain/maintance"
)

func (s *Server) routes() {
	if s.Env.LogHttpRequest {
		s.router.Use(middleware.Logger)
	}
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.CleanPath)
	s.router.Use(middleware.Heartbeat("/ping"))

	// TODO: Pass dbstore
	accounts.LoadRoutes(s.router)
	maintance.LoadRoutes(s.router, s.ServerLogVar)
}
