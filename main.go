package main

import (
	"blue-justice/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")

	router.GET("/", handlers.IndexGet())

	router.Run()
}
