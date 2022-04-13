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

	// Login
	router.GET("/", handlers.LoginRedirect())
	router.GET("/login", handlers.LoginGET())
	router.POST("/login", handlers.LoginPOST())
	// Cadastro
	router.GET("/cadastro", handlers.SignupGet())

	return router
}
