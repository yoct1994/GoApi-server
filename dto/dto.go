package dto

import "time"

type Post struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Name    string    `json:"name"`
	WriteAt time.Time `json:"write_at"`
}
