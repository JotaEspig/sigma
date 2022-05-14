package auth

import (
	"github.com/jmoiron/sqlx"
)

// Gets a user from a database
func GetUser(db *sqlx.DB, username string) (*User, error) {
	u := User{}

	// TODO Jota: Filter results using and array of wanted attributes or whatever

	err := db.Get(&u, "SELECT * FROM \"user\" WHERE username=$1", username)

	return &u, err
}

// Adds an user to a database
func AddUser(db *sqlx.DB, u *User) {
	db.MustExec(
		`INSERT INTO "user"(username, password, email, name)
		VALUES($1, $2, $3, $4)`,
		u.Username, u.HashedPassword, u.Email, u.Name,
	)
}

func RmUser(db *sqlx.DB, username string) {
	db.MustExec(
		`DELETE FROM "user" WHERE username = $1`,
		username,
	)
}
