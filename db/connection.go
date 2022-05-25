package db

import (
	"sigma/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = connInit().getDB()

type Connection struct {
	db *gorm.DB
}

func connInit() *Connection {
	conn := &Connection{}
	conn.connectDB()
	return conn
}

// Connects with a database
func (c *Connection) connectDB() {
	connStr := config.GetDBConfig() // from config.go
	newDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	c.db = newDB
}

// Gets the database variable from the connection
func (c *Connection) getDB() *gorm.DB {
	return c.db
}
