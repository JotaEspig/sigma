package main

import (
	"net/http"
	"os"
	"path/filepath"
	"sigma/handlers"

	"github.com/gin-gonic/gin"
)

// Configures and creates a router
func createRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("static/html/*.html")

	// These lines add a route to every HTML file inside ./html
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Walks inside the folder, checks the filename and then adds as GET
	filepath.Walk(pwd+"/static/html/", func(path string, info os.FileInfo, err error) error {
		if len(info.Name()) < 6 {
			return nil
		}

		idxUntilFileExt := len(info.Name()) - 4
		filePath := "/" + info.Name()
		filePath = filePath[:idxUntilFileExt]

		router.GET(filePath, func(ctx *gin.Context) {
			ctx.HTML(
				http.StatusOK,
				info.Name(),
				nil,
			)
		})
		return nil
	})

	// Loads the css and js folders
	router.Static("css/", "static/css/")
	router.Static("js/", "static/js/")
	router.Static("img/", "static/img/")

	// Login
	router.GET("/", handlers.LoginRedirect())
	router.POST("/login", handlers.LoginPOST())

	// Cadastro
	router.POST("/cadastro", handlers.SignupPOST())

	// Validate User
	router.GET("/validate_user", handlers.ValidateUser())

	return router
}
