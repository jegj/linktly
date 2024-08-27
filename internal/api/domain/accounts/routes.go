package accounts

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/v1/accounts", func(r chi.Router) {
		r.Get(
			"/",
			func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write([]byte("accounts /"))
				if err != nil {
					log.Fatalln("error..")
				}
			},
		)
	})
	return r
}
