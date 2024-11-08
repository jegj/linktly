package links

import (
	"net/http"
	"time"
)

type Link struct {
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
	FolderId    *string    `db:"folder_id" validate:"uuid" json:"folder_id"`
	ExpiresAt   *time.Time `db:"expires_at" validate:"omitempty,expires_at" json:"expires_at"`
	Id          string     `db:"id" json:"id"`
	LinktlyCode string     `db:"linktly_code" validate:"omitempty,http_url" json:"linktly_url"`
	Url         string     `db:"url" validate:"http_url" json:"url"`
	Name        string     `db:"name" validate:"required,min=3,max=255" json:"name"`
	AccountId   string     `db:"account_id" validate:"omitempty,uuid" json:"account_id"`
	Description string     `db:"description" validate:"min=10" json:"description"`
}

type LinkReq struct {
	*Link
}

func (req *LinkReq) Bind(r *http.Request) error {
	return nil
}

type LinkResp struct {
	*Link
}

func (res *LinkResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
