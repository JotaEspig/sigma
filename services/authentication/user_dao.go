package auth

import "database/sql"

// Gets a user from a database
func GetUser(db *sql.DB, username string) (*User, error) {
	var u User

	err := db.QueryRow(
		"SELECT * FROM \"user\" WHERE \"username\" = $1",
		username,
	).Scan(&u.ID, &u.Username, &u.hashedPassword, &u.Email, &u.Name)

	return &u, err
}

// Adds an user to a database
func AddUser(db *sql.DB, u *User) error {
	stmt, err := db.Prepare(
		`INSERT INTO "user"(username, password, email, name)
		VALUES($1, $2, $3, $4)`,
	)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		u.Username, u.hashedPassword, u.Email, u.Name,
	)
	return err
}
