package maintance

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoadRoutes(r chi.Router, serverLogVar *slog.LevelVar) {
	r.Route("/api/maintance", func(r chi.Router) {
		// FIXME: Use Post
		r.Get(
			"/",
			func(w http.ResponseWriter, r *http.Request) {
				serverLogVar.Set(slog.LevelDebug)
				_, err := w.Write([]byte("maintance"))
				if err != nil {
					slog.Error(err.Error())
				}
			},
		)
	})
}
