package db

import "github.com/GoAPI-server/model"

func AutoCreateDB() {
	db := GetDB()

	if !db.Migrator().HasTable(&model.Post{}) {
		db.Migrator().AutoMigrate(&model.Post{})
	}
}
