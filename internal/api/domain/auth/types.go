package auth

import "net/http"

const LinktlyAuthCookieName = "auth_token"

type LoginReq struct {
	Email    string `db:"email" validate:"required,email" json:"email"`
	Password string `db:"password" validate:"required,min=6,max=30,password" json:"password"`
}

func (req *LoginReq) Bind(r *http.Request) error {
	return nil
}

type JwtClaims struct {
	Email string
	Sub   string
	Role  int
}
