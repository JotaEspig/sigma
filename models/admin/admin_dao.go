package admin

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
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

// AutoMigrate the admin table
func init() {
	config.DB.AutoMigrate(&Admin{})
}
