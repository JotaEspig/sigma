package admin

import (
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
}

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
