package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connection contains the database connection
type Connection struct {
	db *gorm.DB
}

// ConnInit initializes a connection with database
func ConnInit() *Connection {
	conn := &Connection{}
	conn.connectDB()
	return conn
}

// Connects with a database
func (c *Connection) connectDB() {
	connStr := getDBConfig()
	newDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	c.db = newDB
}

// GetDB gets the database variable from the connection
func (c *Connection) GetDB() *gorm.DB {
	return c.db
}

// Gets the config to open the database
func getDBConfig() string {
	checkEnv := func(envName, defaultVal string) string {
		if env := os.Getenv(envName); env != "" {
			return env
		}
		return defaultVal
	}

	// Checks if it's running on heroku
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}

	user := checkEnv("DB_USERNAME", "postgres")
	password := checkEnv("DB_PASSWORD", "postgres")
	dbName := checkEnv("DB_DB", "sigma")
	host := checkEnv("DB_HOST", "localhost")
	port := checkEnv("DB_PORT", "5432")

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	return connectionStr
}
