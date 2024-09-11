package accounts

import (
	"net/http"
	"time"

	"github.com/jegj/linktly/internal/api/types"
)

type Account struct {
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	ApiToken  *string    `db:"api_token" json:"api_token"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	Id        string     `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	LastName  string     `db:"lastname" json:"lastname"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"password"`
	Role      int        `db:"role" json:"role"`
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

type GetAccountByIdHandlerReq struct {
	Id string `validate:"required,uuid"`
}

type CreateAccountReq struct {
	Name     string `db:"name" validate:"required,min=3,max=255" json:"name"`
	LastName string `db:"lastname" validate:"required,min=3,max=255" json:"lastname"`
	Email    string `db:"email" validate:"required,email" json:"email"`
	Password string `db:"password" validate:"required,min=6,max=30,password" json:"password"`
}

func (req *CreateAccountReq) Bind(r *http.Request) error {
	return nil
}
