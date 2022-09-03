package user

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"

	"gorm.io/gorm"
)

// Default function to update a user in a database
func UpdateUser(db *gorm.DB, u *User) error {
	return db.Model(u).Omit("username", "password", "type").Updates(u).Error
}

// AutoMigrate the user table
func init() {
	config.DB.AutoMigrate(&User{})
}
