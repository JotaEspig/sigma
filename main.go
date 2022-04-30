package main

import "sigma/services/database"

func main() {
	defer database.Conn.CloseDB()

	// Creates and runs the router
	router := createRouter()
	router.Run()
}
