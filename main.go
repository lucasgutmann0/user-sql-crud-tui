package main

import (
	"fmt"
	"strconv"

	"github.com/gutmanndev/gorm-crud/crud"
	"github.com/gutmanndev/gorm-crud/db/models"
)

func main() {
	var user models.User
	// Create User
	name, role, email, age, password := GetUserInput()
	crud.CreateUser(name, role, email, age, password)

	// Get one user
	user = crud.GetOneUser("name", name)
	fmt.Print("requested user: ")
	fmt.Println(user)

	// Update User
	fmt.Println("Update user data")
	name, role, email, age, password = GetUserInput()
	id := strconv.Itoa(int(user.ID))
	crud.UpdateUser(id, name, role, email, age, password)

	// Get updated user
	user = crud.GetOneUser("id", id)
	fmt.Print("requested user: ")
	fmt.Println(user)

	// Delete User if want to
	crud.DeleteUser(user.ID)

	user = crud.GetOneUser("id", id)
	fmt.Print("requested user: ")
	fmt.Println(user)

}

func GetUserInput() (name string, role string, email string, age int, password string) {

	fmt.Printf("Name: ")
	fmt.Scanln(&name)

	fmt.Printf("Role: ")
	fmt.Scanln(&role)

	fmt.Printf("Email: ")
	fmt.Scanln(&email)

	fmt.Printf("Age: ")
	fmt.Scanln(&age)

	fmt.Printf("Password: ")
	fmt.Scanln(&password)
	return
}
