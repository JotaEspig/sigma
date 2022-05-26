package main

import (
	"net/http"
	"os"
	"path/filepath"
	"sigma/controllers"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// TODO Jota: Set the heroku production and staging

var notToRoute = []string{
	"alunoinfo.html",
}

const relativePathToHTML = "/static/html/"

func setNewRelicMiddleware(router *gin.Engine) {
	nrAppName := os.Getenv("NR_APP_NAME")
	nrAPIKey := os.Getenv("NR_API_KEY")
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(nrAppName),
		newrelic.ConfigLicense(nrAPIKey),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	if err != nil {
		panic(err)
	}

	router.Use(nrgin.Middleware(app))
}

// Gets the type of router engine according to ginMode.
// ginMode should be an env variable
func getRouterEngine() *gin.Engine {
	routerMode := os.Getenv("ROUTER_MODE")
	if routerMode == "release" {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.Use(gin.Recovery())
		// Don't use logs middleware
		return router
	}

	if routerMode == "staging" {
		router := gin.Default()
		setNewRelicMiddleware(router)
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

// Gets the route path according to a file information
func getRoutePath(info os.FileInfo) string {
	idxUntilFileExt := len(info.Name()) - 4
	filePath := "/" + info.Name()
	filePathWithoutExt := filePath[:idxUntilFileExt] // removes the ".html"
	return filePathWithoutExt
}

// Set the routes to a router
func setRoutes(router *gin.Engine) {
	// TODO Jota: Create groups of routes to separate the route paths

	// TODO Jota: Suggestion for the future: remove the filepath walker

	currentWorkingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	absPathToHTML := currentWorkingDir + relativePathToHTML
	// Walks inside the folder, checks the filename and then adds an route for it
	filepath.Walk(absPathToHTML, func(path string, info os.FileInfo, err error) error {
		if len(info.Name()) < 6 {
			return nil
		}
		if !isHTMLToRoute(info.Name(), notToRoute) {
			return nil
		}

		filePathWithoutExt := getRoutePath(info)
		router.GET(filePathWithoutExt, func(ctx *gin.Context) {
			ctx.HTML(
				http.StatusOK, info.Name(), nil,
			)
		})
		return nil
	})

	// Login
	router.GET("/", controllers.LoginRedirect())
	router.POST("/login", controllers.LoginPOST())

	// Cadastro
	router.POST("/cadastro", controllers.SignupPOST())

	// Validates User
	router.GET("/validate/user", controllers.ValidateUser())

	// Get user
	router.GET("/user/:username", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "alunoinfo.html", nil)
	})
	router.POST("/user/:username", controllers.GetUserInfo())
}

func createRouter() *gin.Engine {
	router := getRouterEngine()

	router.LoadHTMLGlob("static/html/*.html")

	// Loads the img, css and js folders
	router.Static("css/", "static/css/")
	router.Static("js/", "static/js/")
	router.Static("img/", "static/img/")

	setRoutes(router)

	return router
}
