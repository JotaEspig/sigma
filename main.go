package main

import "sigma/config"

func main() {
	// Creates superadmin in database if it doesn't exist
	createSuperAdmin(config.DB)

	// Creates and runs the router
	router := createRouter()
	router.Run()
}
