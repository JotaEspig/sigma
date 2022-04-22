package db

import (
	"fmt"
	"os"
)

// Checks if there is a environment variable and return its value,
// if not exists it returns a default value
func checkEnv(envName string, defaultVal string) string {
	if env := os.Getenv(envName); env != "" {
		return env
	}
	return defaultVal
}

func getConfig() (string, string) {
	postgresDriver := "postgres"
	user := checkEnv("DB_USERNAME", "postgres")
	password := checkEnv("DB_PASSWORD", "postgres")
	dbName := checkEnv("DB_DB", "sigma")
	host := checkEnv("DB_HOST", "localhost")
	port := checkEnv("DB_PORT", "5432")

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	return postgresDriver, connectionStr
}
