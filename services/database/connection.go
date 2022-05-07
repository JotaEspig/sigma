package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Conn *Connection

type Connection struct {
	db *sql.DB
}

// Connects with a database
func (c *Connection) connectDB() {
	driver, connStr := getConfig() // from config.go
	newDB, err := sql.Open(driver, connStr)
	if err != nil {
		panic(err.Error())
	}

	if newDB.Ping() != nil {
		panic(err.Error())
	}

	c.db = newDB
}

// Gets the database variable from the connection
func (c *Connection) GetDB() *sql.DB {
	return c.db
}

// Closes the database connection
func (c *Connection) CloseDB() {
	c.db.Close()
}

func init() {
	Conn = &Connection{}
	Conn.connectDB()
}
