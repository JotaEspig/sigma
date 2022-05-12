package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Name     string
	Password string
}

func InitUser(usern, email, name, Password string) *User {
	u := &User{
		Username: usern,
		Email:    email,
		Name:     name,
	}
	passwd, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(passwd)

	return u

}

// Validates the user. It compares the hashed password in the database
// to the password that the user input
func (u *User) Validate(userInput, passInput string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passInput))
	return u.Username == userInput && err == nil
}

// Returns a map containing user info WITHOUT password.
// This map will be send in /validate_user
func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
		"name":     u.Name,
	}
}
