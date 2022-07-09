package main

import (
	"net/http"
	"os"
	"sigma/controllers"
	"sigma/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

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
	// Login
	router.GET("/", controllers.LoginRedirect())
	router.GET("/login", controllers.LoginGET())
	router.POST("/login", controllers.LoginPOST())
	router.GET("/login/validate", controllers.IsLogged())

	// Cadastro
	router.GET("/cadastro", controllers.SignupGET())
	router.POST("/cadastro", controllers.SignupPOST())

	// User group
	user := router.Group("/usuario", middlewares.AuthMiddleware())
	user.GET("", controllers.GetProfilePage())
	user.GET("/get", controllers.GetAllUserInfo())
	user.PUT("/update", controllers.UpdateUser())

	// Public user group (everyone can access this)
	publicUser := user.Group("/:username")
	publicUser.GET("", controllers.GetUserPage())
	publicUser.GET("/get", controllers.GetPublicUserInfo())

	// Student group
	student := router.Group("/aluno", middlewares.IsStudentMiddleware())
	student.GET("", controllers.GetStudentPage())
	student.GET("/get", controllers.GetStudentInfo())

	// Admin group
	admin := router.Group("/admin", middlewares.IsAdminMiddleware())
	admin.GET("", controllers.GetAdminPage())
	admin.GET("/get", controllers.GetAdminInfo())
	admin.PUT("/update", controllers.UpdateAdmin())

	// Admin tools group
	adminTools := admin.Group("/tools")

	// Admin tools to manage others admins
	adminToolsForAdmin := adminTools.Group("/admin/:target",
		middlewares.IsSuperAdminMiddleware())
	adminToolsForAdmin.PUT("/update", controllers.UpdateTargetAdmin())
	adminToolsForAdmin.DELETE("/delete", controllers.DeleteTargetAdmin())
}

func createRouter() *gin.Engine {
	router := getRouterEngine()

	router.LoadHTMLGlob("static/html/*.html")

	// Loads the img, css and js folders
	router.Static("css/", "static/css/")
	router.Static("js/", "static/js/")
	router.Static("img/", "static/img/")

	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})

	setRoutes(router)

	return router
}
