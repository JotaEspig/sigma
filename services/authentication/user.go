package auth

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

//TODO jota: it's needed to separate the struct user to "admin", "teacher", "student"
// Maybe implement something similar to Inheritance

type User struct {
	ID             int
	Username       string
	Name           string
	Surname        string
	Email          string
	HashedPassword string `db:"password"`
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
// This map will be send in /validate_user
func (u *User) ToMap() map[string]interface{} {
	userMap := make(map[string]interface{})
	userMap["id"] = u.ID
	userMap["username"] = u.Username
	userMap["name"] = u.Name
	userMap["surname"] = u.Surname
	userMap["email"] = u.Email
	userMap["type"] = u.Type.String
	return userMap
}
