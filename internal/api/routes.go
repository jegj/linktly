package api

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jegj/linktly/internal/api/domain/accounts"
	"github.com/jegj/linktly/internal/api/domain/auth"
	"github.com/jegj/linktly/internal/api/domain/folders"
	"github.com/jegj/linktly/internal/api/domain/links"
	"github.com/jegj/linktly/internal/api/domain/redirections"
)

func (s *Server) routes() {
	if s.config.LogHttpRequest {
		s.router.Use(middleware.Logger)
	}
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.CleanPath)
	s.router.Use(middleware.Heartbeat("/api/ping"))

	accounts.LoadRoutes(s.router, s.config, s.store)
	auth.LoadRoutes(s.router, s.config, s.store)
	folders.LoadRoutes(s.router, s.config, s.store)
	links.LoadRoutes(s.router, s.config, s.store)
	redirections.LoadRoutes(s.router, s.config, s.store)
	// maintance.LoadRoutes(s.router, s.ServerLogVar)
}
