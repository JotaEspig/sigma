package server

import (
	"sigma/models/admin"
	"sigma/models/user"

	"gorm.io/gorm"
)

// Creates superadmin in database
func createSuperAdmin(db *gorm.DB) {
	// creates defaults values in database
	u := user.InitUser("admin", "admin@gmail.com",
		"admin", "admin", "admin")
	err := db.Create(u).Error
	if err != nil {
		return
	}

	a, err := admin.InitAdmin(u)
	if err != nil {
		panic(err)
	}

	db.Transaction(func(tx *gorm.DB) error {
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
