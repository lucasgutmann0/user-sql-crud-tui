package crud

import (
	"github.com/gutmanndev/gorm-crud/db"
	"github.com/gutmanndev/gorm-crud/db/models"
)

func UpdateUser(
	id string,
	name string,
	role string,
	email string,
	age int,
	password string,
) {
	user := GetOneUser("id", id)

	db.Client.Model(&user).Updates(models.User{
		Name:     name,
		Role:     role,
		Email:    email,
		Age:      age,
		Password: password,
	})
}
