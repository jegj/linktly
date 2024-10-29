package folders

import (
	"net/http"
	"time"
)

type Folder struct {
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
	Id             string     `db:"id" json:"id"`
	Name           string     `db:"name" validate:"required,min=3,max=255" json:"name"`
	ParentFolderId *string    `db:"parent_folder_id" validate:"omitempty,uuid" json:"parent_folder_id"`
	AccountId      string     `db:"account_id" validate:"omitempty,uuid" json:"account_id"`
	Description    string     `db:"description" validate:"min=10" json:"description"`
}

type FolderReq struct {
	*Folder
}

func (req *FolderReq) Bind(r *http.Request) error {
	return nil
}

type FolderResp struct {
	*Folder
}

func (res *FolderResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type FolderDeleteResp struct {
	Id string `json:"id"`
}

func (res *FolderDeleteResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
