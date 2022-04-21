package db

import (
	"fmt"
	"os"
)

const postgresDriver = "postgres"

var (
	user          string
	password      string
	dbName        string
	host          string
	port          string
	connectionStr string
)

func checkEnv(envName string, defaultVal string) string {
	if env := os.Getenv(envName); env != "" {
		return env
	}
	return defaultVal
}

func init() {
	user = checkEnv("DB_USERNAME", "postgres")
	password = checkEnv("DB_PASSWORD", "postgres")
	dbName = checkEnv("DB_DB", "sigma")
	host = checkEnv("DB_HOST", "localhost")
	port = checkEnv("DB_PORT", "5432")
	connectionStr = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}
