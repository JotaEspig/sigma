package main

import (
	"sigma/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")
	router.Static("css/", "css/")
	router.Static("js/", "js/")

	router.GET("/", handlers.IndexGet())

	router.Run()
}
