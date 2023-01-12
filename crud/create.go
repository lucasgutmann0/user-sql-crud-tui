package crud

import (
	"github.com/gutmanndev/gorm-crud/db"
	"github.com/gutmanndev/gorm-crud/db/models"
)

func CreateUser(
	name string,
	Role string,
	email string,
	age int,
	password string) {
	// insert user to db
	db.Client.Create(&models.User{Name: name, Email: email, Age: age, Password: password})
	// db.Create(&User{Code: "D42", Price: 100})
}
