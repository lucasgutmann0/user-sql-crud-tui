package crud

import (
	"github.com/gutmanndev/gorm-crud/db"
	"github.com/gutmanndev/gorm-crud/db/models"
)

func DeleteUser(id uint) {
	var user models.User
	// Delete - delete product
	db.Client.Delete(&user, id)
}
