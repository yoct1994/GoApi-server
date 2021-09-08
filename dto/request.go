package dto

type WritePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

type UpdatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
