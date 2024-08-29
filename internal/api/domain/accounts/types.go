package accounts

import "time"

type Account struct {
	Id        string    `db:"release_date"`
	Name      string    `db:"name"`
	LastName  string    `db:"lastname"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	ApiToken  string    `db:"api_token"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
