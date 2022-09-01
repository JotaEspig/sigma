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
	u.ID = 1
	db.Create(u)

	a, err := admin.InitAdmin(u)
	if err != nil {
		panic(err)
	}

	admin.AddAdmin(db, a)
}
