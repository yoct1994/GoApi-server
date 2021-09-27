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

	err := d.Limit(20).Offset(20 * pageNum).Find(&posts).Error

	return posts, err
}

func UpdatePost(postId int64, updatePost request.UpdatePost) error {
	d := db.GetDB()

	post := model.Post{}

	err := d.Where("id = ?", postId).Find(&post).Error

	post.Title = updatePost.Title
	post.Content = updatePost.Content

	d.Save(&post)

	return err
}
