package db

import (
	"github.com/gutmanndev/gorm-crud/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB = ConnectToDb()

func ConnectToDb() (db *gorm.DB) {
	// Connecion to Db
	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema to sqlite db
	db.AutoMigrate(&models.User{})

	return
}
