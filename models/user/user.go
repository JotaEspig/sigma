package user

import (
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in sigma
type User struct {
	ID             uint   `gorm:"primary_key"`
	Username       string `gorm:"not null;unique"`
	Name           string `gorm:"not null"`
	Surname        string `gorm:"not null"`
	Email          string `gorm:"not null"`
	HashedPassword string `gorm:"not null;column:password"`
	Type           string
}

// InitUser initializes a user struct
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

// Validate validates the user. It compares the hashed password in the database
// to the password that the user input
func (u *User) Validate(username, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return u.Username == username && err == nil
}

// ToMap returns a map containing user info WITHOUT password.
func (u *User) ToMap() map[string]interface{} {
	userMap := make(map[string]interface{})
	userMap["id"] = u.ID
	userMap["username"] = u.Username
	userMap["name"] = u.Name
	userMap["surname"] = u.Surname
	userMap["email"] = u.Email
	userMap["type"] = u.Type
	return userMap
}
