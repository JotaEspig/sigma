package main

import (
	"net/http"
	"sigma/handlers"
	"sigma/services/login"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")

	router.Static("css/", "css/")
	router.Static("js/", "js/")

	// Login
	router.GET("/", handlers.LoginRedirect)
	router.GET("/login", handlers.LoginGET)
	router.POST("/login", handlers.LoginPOST)
	// Cadastro
	router.GET("/cadastro", handlers.SignupGet)

	router.GET("/test", func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if err != nil || token == "" {
			ctx.HTML(
				http.StatusUnauthorized,
				"logintest.html",
				gin.H{
					"ServerResponse": "Você não está logado",
				},
			)
			return
		}

		dtoken, err := login.JWTDefault.ValidateToken(token)
		if err != nil || !dtoken.Valid {
			ctx.SetCookie("auth", "", -1, "", "", false, false)
			ctx.HTML(
				http.StatusUnauthorized,
				"logintest.html",
				gin.H{
					"ServerResponse": "Você não está logado",
				},
			)
			return
		}

		ctx.HTML(
			200,
			"logintest.html",
			nil,
		)
	})

	return router
}
