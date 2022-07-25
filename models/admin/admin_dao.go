package admin

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
	dbPKG "sigma/db"
	"sigma/models/user"

	"gorm.io/gorm"
)

// Slice of all admin params
var AdminParams = []string{
	"id",
	"role",
}

// Slice of public admin params
var PublicAdminParams = []string{
	"id",
	"role",
}

// Adds an admin to a database.
func AddAdmin(db *gorm.DB, a *Admin) error {
	return db.Transaction(func(tx *gorm.DB) error {
		a.User.Type = "admin"
		err := user.UpdateUser(db, a.User)
		if err != nil {
			return err
		}

		err = tx.Create(a).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// Gets an admin from a database
func GetAdmin(db *gorm.DB, username string, params ...string) (*Admin, error) {
	a := &Admin{}

	u, err := user.GetUser(db, username, "id")
	if err != nil {
		return nil, err
	}

	columnsToUse := dbPKG.GetColumns(AdminParams, params...)

	err = db.Select(columnsToUse).Where("id = ?", u.ID).First(a).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(a).Association("User").Find(&a.User)

	return a, err
}

// Updates an admin in a database
func UpdateAdmin(db *gorm.DB, a *Admin) error {
	return db.Model(a).Updates(a).Error
}

// Removes an admin from a database
func RmAdmin(db *gorm.DB, username string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		u, err := user.GetUser(db, username, "id")
		if err != nil {
			return err
		}

		// Updates the type of the user to be empty
		err = db.Model(u).Update("type", "").Error
		if err != nil {
			return err
		}

		return db.Unscoped().Delete(&Admin{}, "id = ?", u.ID).Error
	})
}

// AutoMigrate the admin table
func init() {
	config.DB.AutoMigrate(&Admin{})
}
