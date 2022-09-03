package user

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
)

// AutoMigrate the user table
func init() {
	config.DB.AutoMigrate(&User{})
}
