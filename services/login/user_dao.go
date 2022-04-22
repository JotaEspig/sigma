package login

import "database/sql"

// Gets a user from the database
func GetUser(db *sql.DB, username string) (*User, error) {
	var u User

	err := db.QueryRow(
		"SELECT * FROM \"user\" WHERE \"username\" = $1",
		username,
	).Scan(&u.ID, &u.Username, &u.hashedPassword)

	return &u, err
}
