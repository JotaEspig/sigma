package db

import "database/sql"

var DB *sql.DB = connect(postgresDriver, connectionStr)

func connect(driver, connectionStr string) *sql.DB {
	db, err := sql.Open(driver, connectionStr)
	if err != nil {
		panic(err.Error())
	}

	return db
}
