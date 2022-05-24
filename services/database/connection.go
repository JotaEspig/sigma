package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

func ConnInit() *Connection {
	conn := &Connection{}
	conn.connectDB()
	return conn
}

// Connects with a database
func (c *Connection) connectDB() {
	connStr := getConfig() // from config.go
	newDB, _ := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	c.db = newDB
}

// Gets the database variable from the connection
func (c *Connection) GetDB() *gorm.DB {
	return c.db
}
