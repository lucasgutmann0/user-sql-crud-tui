package crud

import (
	"fmt"

	"github.com/gutmanndev/gorm-crud/db"
	"github.com/gutmanndev/gorm-crud/db/models"
)

func GetAllUser() (users []models.User) {
	db.Client.Find(&users)
	return users
}

func GetOneUser(id string, value string) (user models.User) {
	db.Client.Where(fmt.Sprintf("%s = ?", id), value).First(&user)
	return user
}

func GetMultipleUser(id string, value string) (users models.User) {
	db.Client.Where(fmt.Sprintf("%s = ?", id), value).Find(&users)
	return users
}
