package repository

import (
	"github.com/GoAPI-server/db"
	"github.com/GoAPI-server/dto/request"
	"github.com/GoAPI-server/model"
	"github.com/savsgio/go-logger/v2"
)

func GetPost(postId int64) (model.Post, error) {
	d := db.GetDB()

	var post model.Post

	err := d.First(&post, "id = ?", postId).Error

	return post, err
}

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

	logger.Info(err)

	return posts, err
}

func UpdatePost(postId int64, updatePost request.UpdatePost) error {
	d := db.GetDB()

	post, err := GetPost(postId)

	if err != nil {
		logger.Info(err)
		return err
	}

	if updatePost.Title != "" {
		post.Title = updatePost.Title
	}

	if updatePost.Content != "" {
		post.Content = updatePost.Content
	}

	err = d.Save(&post).Error
	if err != nil {
		logger.Info(err)
		return err
	}

	return err
}

func DeletePost(postId int64) error {
	d := db.GetDB()

	post, err := GetPost(postId)

	if err != nil {
		logger.Info(err)
		return err
	}

	err = d.Delete(&post).Error

	if err != nil {
		logger.Info(err)
		return err
	}

	return err
}
