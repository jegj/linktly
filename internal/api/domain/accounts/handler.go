package accounts

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/jegj/linktly/internal/api/types"
)

type AccountHandler struct {
	service AccountService
}

func (s AccountHandler) GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Debug message accoutns")
	slog.Info("Info message accoutns")

	account, error := s.service.GetAccountById("0191b574-3e8f-7e79-9848-4f9bb436e4b2")

	if error != nil {
		err := render.Render(w, r, types.NewLinktlyError(error, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)))
		if err != nil {
			slog.Error(err.Error())
		}
	} else {
		resp := &AccountResp{
			Account: account,
		}
		render.Status(r, http.StatusOK)
		err := render.Render(w, r, resp)
		if err != nil {
			slog.Error(err.Error())
		}
	}
}
