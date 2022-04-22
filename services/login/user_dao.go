package login

import "database/sql"

// Gets a user from the database
func GetUser(db *sql.DB, username string) (*User, error) {
	u := &User{}

	err := db.QueryRow(
		"SELECT * FROM \"user\" WHERE \"username\" = $1",
		username,
	).Scan(u.ID, u.Username, u.hashedPassword)

	return u, err
}

func AddUser(db *sql.DB, u *User) error {
	stmt, err := db.Prepare(`INSERT INTO "user"(username, password)
		VALUES($1, $2)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Username, u.hashedPassword)
	return err
}
