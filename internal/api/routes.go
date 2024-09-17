package api

import (
	"context"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/jegj/linktly/internal/api/domain/accounts"
	"github.com/jegj/linktly/internal/api/domain/auth"
)

func (s *Server) routes(ctx context.Context) {
	if s.config.LogHttpRequest {
		s.router.Use(middleware.Logger)
	}
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.CleanPath)
	s.router.Use(middleware.Heartbeat("/ping"))

	accounts.LoadRoutes(ctx, s.router, s.config, s.store)
	auth.LoadRoutes(ctx, s.router, s.config, s.store)
	// maintance.LoadRoutes(s.router, s.ServerLogVar)
}
