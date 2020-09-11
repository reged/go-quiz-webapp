package models

// WikiPage describes simple wiki page
type WikiPage struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  []byte `json:"body"`
}
