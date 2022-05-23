package userauth

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Adds an user to a database.
// Panics if something goes wrong.
func AddUser(db *sqlx.DB, u *User) {
	db.MustExec(
		`INSERT INTO "user"(username, password, name, surname, email)
		VALUES($1, $2, $3, $4, $5)`,
		u.Username, u.HashedPassword, u.Name, u.Surname, u.Email,
	)
}

// Gets an user from a database
func GetUser(db *sqlx.DB, username string, columns ...string) (*User, error) {
	u := User{}

	columnsStr := strings.Join(columns, ",")
	// If there is no arguments, sets it to default
	if len(columns) == 0 {
		columnsStr = "*"
	}

	// TODO Jota: Implement something to "clean" columnsStr
	// the reason is to avoid SQLInjection (if that's possible)
	sqlQuery := fmt.Sprintf("SELECT %s FROM \"user\" WHERE username=$1", columnsStr)

	err := db.Get(&u, sqlQuery, username)

	return &u, err
}

// Removes an user
func RmUser(db *sqlx.DB, username string) {
	db.MustExec(
		`DELETE FROM "user" WHERE username = $1`,
		username,
	)
}
