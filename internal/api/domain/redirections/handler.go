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
	link, err := l.service.GetLinkByCode(r.Context(), code)
	if err != nil {
		// TODO: HANDLE THIS WITH A REVER PROXY
		http.Redirect(w, r, "/404.html", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, link.Url, http.StatusMovedPermanently)
	}
}
