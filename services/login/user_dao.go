package login

import "sigma/services/db"

// Gets a user from the database
func GetUser(username string) (*User, error) {
	var u User

	err := db.DB.QueryRow(
		"SELECT * FROM \"user\" WHERE \"username\" = $1",
		username,
	).Scan(&u.ID, &u.Username, &u.hashedPassword)

	return &u, err
}