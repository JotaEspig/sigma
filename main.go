package main

import "sigma/handlers"

func main() {
	defer handlers.Conn.CloseDB()
	// Creates and runs the router
	router := createRouter()
	router.Run()
}
