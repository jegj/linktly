package accounts

import (
	"net/http"

	"github.com/jegj/linktly/internal/api/types"
)

type AccountResp struct {
	*Account
	Password types.Omit `json:"password,omitempty"`
}

func (rd *AccountResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
