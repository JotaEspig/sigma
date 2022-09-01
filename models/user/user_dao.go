package user

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"

	"gorm.io/gorm"
)

// Slice of all user params
var UserParams = []string{
	"id",
	"username",
	"name",
	"surname",
	"email",
	"password",
	"type",
}

// Slice of public user params
var PublicUserParams = []string{
	"id",
	"username",
	"name",
	"surname",
	"email",
	"type",
}

// Default function to update a user in a database
func UpdateUser(db *gorm.DB, u *User) error {
	return db.Model(u).Omit("username", "password", "type").Updates(u).Error
}

// Removes a user from a database
func RmUser(db *gorm.DB, username string) error {
	return db.Unscoped().Delete(&User{}, "username = ?", username).Error
}

// AutoMigrate the user table
func init() {
	config.DB.AutoMigrate(&User{})
}
