package main

import (
	"sigma/handlers"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")
	router.Static("css/", "css/")
	router.Static("js/", "js/")

	router.GET("/", handlers.LoginRedirect())
	router.GET("/login", handlers.LoginGet())

	return router
}
