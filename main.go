package main

import "sigma/services/db"

func main() {
	defer db.DB.Close()

	router := createRouter()
	router.Run()
}
