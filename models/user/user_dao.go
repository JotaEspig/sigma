package user

import (
	"gorm.io/gorm"
)

func GetColumns(columns ...string) interface{} {
	var columnsToUse interface{}

	columnsToUse = "*"
	if len(columns) != 0 {
		columnsToUse = columns
	}

	return columnsToUse
}

// Adds an user to a database.
func AddUser(db *gorm.DB, u *User) {
	db.Create(u)
}

// Gets an user from a database
func GetUser(db *gorm.DB, username string, columns ...string) *User {
	u := User{}

	columnsToUse := GetColumns(columns...)

	db.Select(columnsToUse).Where("username = ?", username).First(&u)

	return &u
}

// Removes an user
func RmUser(db *gorm.DB, username string) {
	db.Where("username = ?", username).Delete(&User{})
}
