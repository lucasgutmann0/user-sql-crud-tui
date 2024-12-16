package models

type User struct {
	ID       uint   `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Role     string `json:"role" db:"role"`
	Email    string `json:"email" db:"email"`
	Age      int    `json:"age" db:"age"`
	Password string `json:"password" db:"password"`
}
