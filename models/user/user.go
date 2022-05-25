package user

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//TODO jota: it's needed to separate the struct user to "admin", "teacher", "student"
// Maybe implement something similar to Inheritance

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	Name           string `gorm:"not null"`
	Surname        string `gorm:"not null"`
	Email          string `gorm:"not null"`
	HashedPassword string `gorm:"not null;column:password"`
	Type           sql.NullString
}

func InitUser(usern, email, name, surname, password string) *User {
	u := &User{
		Username: usern,
		Name:     name,
		Surname:  surname,
		Email:    email,
	}
	hashedPasswd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.HashedPassword = string(hashedPasswd)

	return u
}

// Validates the user. It compares the hashed password in the database
// to the password that the user input
func (u *User) Validate(userInput, passInput string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(passInput))
	return u.Username == userInput && err == nil
}

// Returns a map containing user info WITHOUT password.
func (u *User) ToMap() map[string]interface{} {
	userMap := make(map[string]interface{})
	userMap["id"] = u.Model.ID
	userMap["username"] = u.Username
	userMap["name"] = u.Name
	userMap["surname"] = u.Surname
	userMap["email"] = u.Email
	userMap["type"] = u.Type.String
	return userMap
}
