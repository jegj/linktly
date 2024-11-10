package redirections

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RedirectionsHandler struct {
	service RedirectionsService
}

func (l RedirectionsHandler) GetLinkByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	// TODO: VALIDATE CODE

	link, err := l.service.GetLinkByCode(r.Context(), code)
	if err != nil {
		// TODO: MOVE TO 404 PAGE
		http.Redirect(w, r, link.Url, http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, link.Url, http.StatusMovedPermanently)
	}
}
