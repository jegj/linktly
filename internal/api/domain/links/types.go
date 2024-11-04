package links

import (
	"net/http"
	"time"
)

type Link struct {
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
	Id          string     `db:"id" json:"id"`
	LinktlyUrl  string     `db:"linktly_url" json:"linktly_url"`
	Url         string     `db:"url" json:"url"`
	Name        string     `db:"name" validate:"required,min=3,max=255" json:"name"`
	FolderId    *string    `db:"parent_folder_id" validate:"omitempty,uuid" json:"parent_folder_id"`
	AccountId   string     `db:"account_id" validate:"omitempty,uuid" json:"account_id"`
	Description string     `db:"description" validate:"min=10" json:"description"`
	ExpiresAt   *time.Time `db:"expires_at" json:"expires_at"`
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
