package user

import (
	dbPKG "sigma/db"

	"gorm.io/gorm"
)

// Adds an user to a database.
func AddUser(db *gorm.DB, u *User) error {
	return db.Create(u).Error
}

// Gets an user from a database
func GetUser(db *gorm.DB, username string, columns ...string) (*User, error) {
	u := &User{}

	columnsToUse := dbPKG.GetColumns(columns...)

	err := db.Select(columnsToUse).Where("username = ?", username).First(u).Error

	return u, err
}

// Removes an user from a database
func RmUser(db *gorm.DB, username string) error {
	return db.Unscoped().Where("username = ?", username).Delete(&User{}).Error
}
