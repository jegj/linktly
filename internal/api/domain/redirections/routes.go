package redirections

import (
	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/config"
	"github.com/jegj/linktly/internal/store"
)

func LoadRoutes(r chi.Router, config config.Config, store *store.PostgresStore) {
	redirectionsRepository := GetNewRlinkRepository(store)
	redirectionsService := RedirectionsService{
		Repository: redirectionsRepository,
	}
	redirectionsHandler := RedirectionsHandler{
		service: redirectionsService,
	}

	r.Route("/api/v1/r", func(r chi.Router) {
		r.Get("/{code}", redirectionsHandler.GetLinkByCode)
	})
}
