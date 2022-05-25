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
func AddUser(db *gorm.DB, u *User) error {
	return db.Create(u).Error
}

// Gets an user from a database
func GetUser(db *gorm.DB, username string, columns ...string) (*User, error) {
	u := User{}

	columnsToUse := GetColumns(columns...)

	err := db.Select(columnsToUse).Where("username = ?", username).First(&u).Error

	return &u, err
}

// Removes an user
func RmUser(db *gorm.DB, username string) {
	db.Where("username = ?", username).Delete(&User{})
}
