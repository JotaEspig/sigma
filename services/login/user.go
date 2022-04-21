package login

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id             int
	username       string
	hashedPassword string
}

func DefaultUser() *User {
	return &User{
		username:       "teste",
		hashedPassword: "$2a$10$CsTxuGv/5Y7KUl65AdspPeT1jMjpJePt6Hoi9uKGrsWt3mVdSZK/W",
	}
}

func (u *User) Validate(userInput, passInput string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.hashedPassword), []byte(passInput))
	println(len(u.hashedPassword))
	return u.username == userInput && err == nil
}

/*password := []byte("MyDarkSecret")

// Hashing the password with the default cost of 10
hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
if err != nil {
	panic(err)
}
fmt.Println(string(hashedPassword))

// Comparing the password with the hash
err = bcrypt.CompareHashAndPassword(hashedPassword, password)
fmt.Println(err) // nil means it is a match
*/
