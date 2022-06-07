package main

import (
	"os"
	"sigma/controllers"
	"sigma/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// TODO Jota: Set the heroku production and staging

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

// Set the routes to a router
func setRoutes(router *gin.Engine) {
	// TODO Jota: Create groups of routes to separate the route paths

	// Login
	router.GET("/", controllers.LoginRedirect())
	router.GET("/login", controllers.LoginGET())
	router.POST("/login", controllers.LoginPOST())

	// Cadastro
	router.GET("/cadastro", controllers.SignupGET())
	router.POST("/cadastro", controllers.SignupPOST())

	user := router.Group("/user")

	// Get user
	user.GET("/:username", controllers.GetUserInfoPage())
	user.GET("/:username/get", controllers.GetPublicUserInfo())

	// TODO Jota: Check if validate is logical
	// Validates User
	user.GET("/validate", controllers.ValidateUser())
	user.GET("/:username/validate",
		middlewares.AuthMiddleware(), controllers.GetAllUserInfo())

	// Aluno
	user.GET("/:username/aluno",
		middlewares.AuthMiddleware(), controllers.GetAlunoPage())
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
