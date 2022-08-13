package main

import (
	"sigma/server"
)

func main() {
	// Creates and runs the router
	router := server.CreateRouter()
	router.Run()
}
