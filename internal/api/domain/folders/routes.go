package folders

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/api/handlers"
	"github.com/jegj/linktly/internal/api/jwt"
	"github.com/jegj/linktly/internal/config"
	"github.com/jegj/linktly/internal/store"
)

func LoadRoutes(r chi.Router, config config.Config, store *store.PostgresStore) {
	folderRepository := GetNewFolderRepository(store)
	folderService := FolderService{
		Repository: folderRepository,
	}

	folderHandler := FolderHandler{
		service: folderService,
	}
	publicKey, error := config.GetPublicKey()
	if error != nil {
		slog.Error(error.Error())
	} else {
		r.Route("/api/v1/folders", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(jwt.AuthMiddleware(*publicKey))
				r.Method("POST", "/", handlers.CentralizedErrorHandler(folderHandler.CreateFolder))
				r.Method("GET", "/", handlers.CentralizedErrorHandler(folderHandler.GetFoldersByUserId))
				r.Method("DELETE", "/{id}", handlers.CentralizedErrorHandler(folderHandler.DeleteFoldersByIdAndUserId))
				r.Method("PATCH", "/{id}", handlers.CentralizedErrorHandler(folderHandler.PatchFoldersByIdAndUserId))
			})
		})
	}
}
