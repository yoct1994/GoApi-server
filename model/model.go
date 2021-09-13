package model

import "time"

type Post struct {
	Id      int64 `gorm:"primaryKey;autoIncrement"`
	Title   string
	Content string
	Name    string
	WriteAt time.Time `gorm:"autoCreateTime"`
}
