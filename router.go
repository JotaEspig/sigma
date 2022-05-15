package main

import (
	"net/http"
	"os"
	"path/filepath"
	"sigma/handlers"

	"github.com/gin-gonic/gin"
)

// Gets the type of router engine according to ginMode.
// ginMode should be an env variable
func getRouterEngine(ginMode string) *gin.Engine {
	if ginMode == "release" {
		router := gin.New()
		router.Use(gin.Recovery())
		// Don't use logs middleware
		return router
	}

	return gin.Default()
}

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
	router := getRouterEngine(os.Getenv("GIN_MODE"))

	router.LoadHTMLGlob("static/html/*.html")

	// These lines add a route to every HTML file inside ./html (with exceptions)
	notToRoute := []string{
		"alunoinfo.html",
	}

	currentWorkingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	pathToHTML := currentWorkingDir + "/static/html/"
	// Walks inside the folder, checks the filename and then adds an route for it
	filepath.Walk(pathToHTML, func(path string, info os.FileInfo, err error) error {
		if len(info.Name()) < 6 {
			return nil
		}
		if !isHTMLToRoute(info.Name(), notToRoute) {
			return nil
		}

		idxUntilFileExt := len(info.Name()) - 4
		filePath := "/" + info.Name()
		filePathWithoutExt := filePath[:idxUntilFileExt] // removes the ".html"

		router.GET(filePathWithoutExt, func(ctx *gin.Context) {
			ctx.HTML(
				http.StatusOK, info.Name(), nil,
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
	router.GET("/getloggeduser", handlers.GetLoggedUserInfo())

	router.GET("/getuser", handlers.GetUserInfo())

	return router
}
