package auth

import (
	"golang.org/x/crypto/bcrypt"
)

//TODO jota: it's needed to separate the struct user to "admin", "teacher", "student"
// Maybe implement something similar to Inheritance

type User struct {
	ID             int
	Username       string
	Email          string
	Name           string
	HashedPassword string `db:"password"`
}

func InitUser(usern, email, name, password string) *User {
	u := &User{
		Username: usern,
		Email:    email,
		Name:     name,
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
	userMap["email"] = u.Email
	return userMap
}
