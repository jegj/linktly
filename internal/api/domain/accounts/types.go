package accounts

import (
	"net/http"
	"time"

	"github.com/jegj/linktly/internal/api/types"
)

type Account struct {
	Id        string     `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	LastName  string     `db:"lastname" json:"lastname"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"password"`
	ApiToken  *string    `db:"api_token" json:"api_token"`
	Role      int        `db:"role" json:"role"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type AccountResp struct {
	*Account
	Password types.Omit `json:"password,omitempty"`
	ApiToken types.Omit `json:"api_token,omitempty"`
}

func (res *AccountResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type AccountReq struct {
	*Account
}

func (req *AccountReq) Bind(r *http.Request) error {
	return nil
}
