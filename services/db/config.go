package db

import (
	"fmt"
	"os"
)

const PostgresDriver = "postgres"

var (
	User          string
	Password      string
	DbName        string
	Host          string
	Port          string
	ConnectionStr string
)

func checkEnv(envName string, defaultVal string) string {
	if env := os.Getenv(envName); env != "" {
		return env
	}
	return defaultVal
}

func init() {
	User = checkEnv("DB_USERNAME", "postgres")
	Password = checkEnv("DB_PASSWORD", "postgres")
	DbName = checkEnv("DB_DB", "sigma")
	Host = checkEnv("DB_HOST", "localhost")
	Port = checkEnv("DB_PORT", "5432")
	ConnectionStr = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
}
