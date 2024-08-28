package accounts

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoadRoutes(r chi.Router) {
	r.Route("/api/v1/accounts", func(r chi.Router) {
		r.Get(
			"/",
			func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write([]byte("accounts"))
				if err != nil {
					log.Fatal(err.Error())
				}
			},
		)

		r.Get(
			"/info",
			func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write([]byte("accounts info"))
				if err != nil {
					log.Fatal(err.Error())
				}
			},
		)
	})
}
