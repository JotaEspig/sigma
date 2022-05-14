package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Creates and runs the router
	router := createRouter()
	router.Run()
}
