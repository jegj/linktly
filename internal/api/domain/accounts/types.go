package accounts

import (
	"time"
)

type Account struct {
	Id        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	LastName  string    `db:"lastname" json:"lastname"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	ApiToken  string    `db:"api_token" json:"api_token"`
	Role      string    `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
