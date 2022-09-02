package admin

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
	dbPKG "sigma/db"
	"sigma/models/user"

	"gorm.io/gorm"
)

// AdminParams is a slice of all admin params
var AdminParams = []string{
	"id",
	"role",
}

// PublicAdminParams is a slice of public admin params
var PublicAdminParams = []string{
	"id",
	"role",
}

// AddAdmin adds an admin to a database.
func AddAdmin(db *gorm.DB, a *Admin) error {
	return db.Transaction(func(tx *gorm.DB) error {
		a.User.Type = "admin"
		err := db.Model(a.User).Update("type", a.User.Type).Error
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

// GetAdmin gets an admin from a database
func GetAdmin(db *gorm.DB, username string, params ...string) (*Admin, error) {
	a := &Admin{}

	u := user.User{}
	err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
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

// UpdateAdmin is the default function to update an admin in a database
func UpdateAdmin(db *gorm.DB, a *Admin) error {
	return db.Model(a).Omit("id").Updates(a).Error
}

// RmAdmin removes an admin from a database
func RmAdmin(db *gorm.DB, username string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		u := user.User{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
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
