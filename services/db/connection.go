package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB = connect()

// Connects with a database
func connect() *sql.DB {
	driver, connStr := getConfig()
	db, err := sql.Open(driver, connStr)
	if err != nil {
		panic(err.Error())
	}

	return db
}