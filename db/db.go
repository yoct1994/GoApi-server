package db

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDB() error {
	var dbConfig map[string]string
	dbConfig, err := godotenv.Read()

	dnsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["name"],
	)

	db, err := gorm.Open(mysql.Open(dnsn), &gorm.Config{})

	if err != nil {
		logger.Infof("Failed to open database")
		return err
	}

	database = db
	logger.Infof("Database success connecting")

	return err
}

func GetDB() *gorm.DB {
	return database
}

func CloseDB() {
	db, _ := database.DB()
	db.Close()
}
