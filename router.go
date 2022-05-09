package main

import (
	"sigma/handlers"

	"github.com/gin-gonic/gin"
)

// Configures and creates a router
func createRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")

	// Loads the css and js folders
	router.Static("css/", "css/")
	router.Static("js/", "js/")

	// Login
	router.GET("/", handlers.LoginRedirect())
	router.GET("/login", handlers.LoginGET())
	router.POST("/login", handlers.LoginPOST())

	// Cadastro
	router.GET("/cadastro", handlers.SignupGET())
	router.POST("/cadastro", handlers.SignupPOST())

	// Validate User
	router.GET("/validate_user", handlers.ValidateUser())

	router.GET("/aluno", func(ctx *gin.Context) {
		ctx.HTML(
			200,
			"aluno.html",
			nil,
		)
	})

	return router
}
