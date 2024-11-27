package links

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/api/handlers"
	"github.com/jegj/linktly/internal/api/jwt"
	"github.com/jegj/linktly/internal/config"
	"github.com/jegj/linktly/internal/store"
)

func LoadRoutes(r chi.Router, config config.Config, store *store.PostgresStore) {
	linkRepository := GetNewLinkRepository(store)
	linkService := LinksService{
		Repository: linkRepository,
	}
	linkHandler := LinksHandler{
		service: linkService,
	}
	publicKey, error := config.GetPublicKey()
	if error != nil {
		slog.Error(error.Error())
	} else {
		r.Route("/api/v1/links", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(jwt.AuthMiddleware(*publicKey))
				r.Method("POST", "/", handlers.CentralizedErrorHandler(linkHandler.CreateLink))
				r.Method("GET", "/{id}", handlers.CentralizedErrorHandler(linkHandler.GetLink))
				// r.Method("GET", "/{folderId}/links/{id}", handlers.CentralizedErrorHandler(folderHandler.GetFolderByIdAndUserId))
				// r.Method("DELETE", "/{folderId}/links/{id}", handlers.CentralizedErrorHandler(folderHandler.DeleteFoldersByIdAndUserId))
				// r.Method("PATCH", "/{folderId}/links/{id}", handlers.CentralizedErrorHandler(folderHandler.PatchFoldersByIdAndUserId))
			})
		})
	}
}
