package dto

import "time"

type PostResponse struct {
	Title   string    `json:"title"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
	WriteAt time.Time `json:"write_at"`
}

type ErrorRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
