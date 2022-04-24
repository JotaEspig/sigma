package main

import "sigma/services/db"

func main() {
	// Configures and connects to the database
	db.DB = db.Connect()
	defer db.DB.Close()

	// Creates and runs the router
	router := createRouter()
	router.Run()
}
