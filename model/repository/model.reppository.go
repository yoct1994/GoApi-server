package repository

import (
	"github.com/GoAPI-server/db"
	"github.com/GoAPI-server/dto/request"
	"github.com/GoAPI-server/model"
)

func CreatePost(post request.WritePost) error {
	d := db.GetDB()

	data := model.Post{
		Title:   post.Title,
		Content: post.Content,
		Name:    post.Name,
	}

	err := d.Create(&data).Error

	return err
}

func GetAllPosts(pageNum int) ([]*model.Post, error) {
	d := db.GetDB()

	posts := make([]*model.Post, 0)

	err := d.Find(&posts).Limit(20).Offset(20 * pageNum).Error

	return posts, err
}
