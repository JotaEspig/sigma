package user

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	dbPKG "sigma/db"

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

// Adds a user to a database.
func AddUser(db *gorm.DB, u *User) error {
	return db.Create(u).Error
}

// Gets a user from a database
func GetUser(db *gorm.DB, username string, params ...string) (*User, error) {
	u := &User{}

	columnsToUse := dbPKG.GetColumns(UserParams, params...)

	err := db.Select(columnsToUse).Where("username = ?", username).First(u).Error

	return u, err
}

// Updates a user in a database
func UpdateUser(db *gorm.DB, u *User) error {
	return db.Model(u).Updates(u).Error
}

// Removes a user from a database
func RmUser(db *gorm.DB, username string) error {
	return db.Unscoped().Delete(&User{}, "username = ?", username).Error
}

// AutoMigrate the user table
func init() {
	dbPKG.DB.AutoMigrate(&User{})
}
