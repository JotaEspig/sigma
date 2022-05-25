package config

import (
	"fmt"
	"os"
	"sigma/auth"
)

// JWTService variable
var DefaultJWT = auth.JWTAuthService()

// Checks if there is a environment variable and return its value,
// if not exists it returns a default value
func checkEnv(envName string, defaultVal string) string {
	if env := os.Getenv(envName); env != "" {
		return env
	}
	return defaultVal
}

// Gets the config to open the database
func GetDBConfig() string {
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
