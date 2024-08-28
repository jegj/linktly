package accounts

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoadRoutes(r chi.Router) {
	r.Route("/api/v1/accounts", func(r chi.Router) {
		r.Get(
			"/",
			func(w http.ResponseWriter, r *http.Request) {
				slog.Debug("Debug message accoutns")
				slog.Info("Info message accoutns")

				_, err := w.Write([]byte("accounts"))
				if err != nil {
					slog.Error(err.Error())
				}
			},
		)

		r.Get(
			"/info",
			func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write([]byte("accounts info"))
				if err != nil {
					slog.Error(err.Error())
				}
			},
		)
	})
}
