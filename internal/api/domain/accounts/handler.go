package accounts

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Debug message accoutns")
	slog.Info("Info message accoutns")

	account := &Account{
		Id:       "10101",
		Name:     "Javier",
		LastName: "Galarza",
		Email:    "jegj57@gmail.com",
	}

	resp := &AccountResp{
		Account: account,
	}
	render.Status(r, http.StatusOK)
	err := render.Render(w, r, resp)
	if err != nil {
		slog.Error(err.Error())
	}
}
