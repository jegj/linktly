package redirections

type Rlink struct {
	Url string `db:"url" validate:"http_url" json:"url"`
}
