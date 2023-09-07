package api

type Category struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	ParentID int        `json:"parentid"`
	Children []Category `json:children"`
}
