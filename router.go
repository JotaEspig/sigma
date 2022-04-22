package main

import (
	"net/http"
	"sigma/handlers"
	"sigma/services/login"

	"github.com/dgrijalva/jwt-go"
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
	router.GET("/cadastro", handlers.SignupGet())

	router.GET("/test", func(ctx *gin.Context) {
		ctx.HTML(
			200,
			"logintest.html",
			nil,
		)
	})

	router.POST("/test", func(ctx *gin.Context) {
		resp := struct {
			Token string `json:"token"`
		}{}
		ctx.BindJSON(&resp)
		if resp.Token == "" {
			ctx.JSON(
				http.StatusUnauthorized,
				nil,
			)
			return
		}

		dtoken, err := login.JWTDefault.ValidateToken(resp.Token)
		if err != nil || !dtoken.Valid {
			ctx.JSON(
				http.StatusUnauthorized,
				nil,
			)
			return
		}

		claims := dtoken.Claims.(jwt.MapClaims)

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"username": claims["username"],
			},
		)
	})

	return router
}
