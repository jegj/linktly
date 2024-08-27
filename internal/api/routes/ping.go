package route

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PingRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("pong live reloaded"))
			if err != nil {
				log.Fatalln("error..")
			}
		},
	)
	return r
}
