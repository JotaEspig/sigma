package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection struct {
	db *sqlx.DB
}

func ConnInit() *Connection {
	conn := &Connection{}
	conn.connectDB()
	return conn
}

// Connects with a database
func (c *Connection) connectDB() {
	driver, connStr := getConfig() // from config.go
	newDB, err := sqlx.Connect(driver, connStr)
	if err != nil {
		panic(err.Error())
	}

	if newDB.Ping() != nil {
		panic(err.Error())
	}

	c.db = newDB
}

// Gets the database variable from the connection
func (c *Connection) GetDB() *sqlx.DB {
	return c.db
}

// Closes the database connection
func (c *Connection) CloseDB() {
	c.db.Close()
}
