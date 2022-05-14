package main

import (
	"net/http"
	"os"
	"path/filepath"
	"sigma/handlers"

	"github.com/gin-gonic/gin"
)

// Checks if file needs to be route
// It needed by default, and it's not when the filename is in the notToRoute
func isHTMLToRoute(filename string, notToRoute []string) bool {
	for _, val := range notToRoute {
		if val == filename {
			return false
		}
	}
	return true
}

// Configures and creates a router
func createRouter() *gin.Engine {
	var router *gin.Engine

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		router = gin.New()
		router.Use(gin.Recovery())
	} else {
		router = gin.Default()
	}

	router.LoadHTMLGlob("static/html/*.html")

	// These lines add a route to every HTML file inside ./html
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	notToRoute := []string{
		"alunoinfo.html",
	}
	// Walks inside the folder, checks the filename and then adds as GET
	filepath.Walk(pwd+"/static/html/", func(path string, info os.FileInfo, err error) error {
		if len(info.Name()) < 6 {
			return nil
		}
		if !isHTMLToRoute(info.Name(), notToRoute) {
			return nil
		}

		idxUntilFileExt := len(info.Name()) - 4
		filePath := "/" + info.Name()
		filePathWithoutExt := filePath[:idxUntilFileExt]

		router.GET(filePathWithoutExt, func(ctx *gin.Context) {
			ctx.HTML(
				http.StatusOK,
				info.Name(),
				nil,
			)
		})
		return nil
	})

	// Loads the img, css and js folders
	router.Static("css/", "static/css/")
	router.Static("js/", "static/js/")
	router.Static("img/", "static/img/")

	// Login
	router.GET("/", handlers.LoginRedirect())
	router.POST("/login", handlers.LoginPOST())

	// Cadastro
	router.POST("/cadastro", handlers.SignupPOST())

	// Validate User
	router.GET("/validateuser", handlers.GetUserInfo())

	return router
}
