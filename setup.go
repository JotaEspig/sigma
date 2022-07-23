package main

import (
	"sigma/models/admin"
	"sigma/models/user"

	"gorm.io/gorm"
)

// Creates superadmin in database
func createSuperAdmin(db *gorm.DB) {
	// creates defaults values in database
	u := user.InitUser("SUPERADMIN", "superadmin@gmail.com",
		"SUPERADMIN", "SUPERADMIN", "SUPERADMIN")
	u.ID = 1
	user.AddUser(db, u)

	a, err := admin.InitAdmin(u)
	if err != nil {
		panic(err)
	}

	admin.AddAdmin(db, a)
}
