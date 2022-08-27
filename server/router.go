package server

import (
	"io"
	"net/http"
	"os"
	"sigma/config"
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
		gin.DisableConsoleColor()

		// Logging to a file.
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)

		return gin.Default()
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
	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})

	// Login
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/login")
	})
	router.GET("/login", controllers.GetLoginPage())
	router.POST("/login", controllers.Login())
	router.GET("/login/validate", controllers.IsLogged())
	router.GET("/logout", controllers.Logout())

	// Cadastro
	router.GET("/cadastro", controllers.SignupGET())
	router.POST("/cadastro", controllers.SignupPOST())

	router.GET("/search/users/:username", controllers.SearchUsers())

	// User group
	user := router.Group("/usuario", middlewares.AuthMiddleware())
	user.GET("", controllers.GetProfilePage())
	user.GET("/get", controllers.GetAllUserInfo())
	user.PUT("/update", controllers.UpdateUser())

	// Public user group (everyone can access this)
	publicUser := router.Group("/:username")
	publicUser.GET("", controllers.GetUserPage())
	publicUser.GET("/get", controllers.GetPublicUserInfo())

	// Student group
	student := router.Group("/aluno", middlewares.IsStudentMiddleware())
	student.GET("", controllers.GetStudentPage())
	student.GET("/get", controllers.GetStudentInfo())

	// Teacher group
	teacher := router.Group("/professor", middlewares.IsTeacherMiddleware())
	teacher.GET("", controllers.GetTeacherPage())
	teacher.GET("/get", controllers.GetTeacherInfo())
	teacher.GET("/update", controllers.UpdateTeacher())

	// Admin group
	admin := router.Group("/admin", middlewares.IsAdminMiddleware())
	admin.GET("", controllers.GetAdminPage())
	admin.GET("/get", controllers.GetAdminInfo())
	admin.PUT("/update", controllers.UpdateAdmin())

	// Admin tools group
	adminTools := admin.Group("/tools")

	// Admin tools to manage classrooms
	adminToolsForClassroom := adminTools.Group("/classroom")
	adminToolsForClassroom.POST("/add", controllers.AddClassroom())
	adminToolsForClassroom.GET("/get", controllers.GetAllClassroomsInfo())
	adminToolsForClassroom.GET("/:id/get", controllers.GetClassroomInfo())

	// Admin tools to manage others admins
	adminToolsForAdmin := adminTools.Group("/admin",
		middlewares.IsSuperAdminMiddleware())
	adminToolsForAdmin.POST("/add", controllers.AddAdmin())
	adminToolsForAdmin.GET("/:username/get", controllers.GetAdminInfo())
	adminToolsForAdmin.PUT("/:username/update", controllers.UpdateAdmin())
	adminToolsForAdmin.DELETE("/:username/delete", controllers.DeleteAdmin())
}

func CreateRouter() *gin.Engine {
	router := getRouterEngine()

	router.LoadHTMLGlob("static/html/**/*.html")

	// Loads the img, css and js folders
	router.Static("css/", "static/css/")
	router.Static("js/", "static/js/")
	router.Static("img/", "static/img/")

	setRoutes(router)
	createSuperAdmin(config.DB)
	return router
}

func CreateTestRouter() *gin.Engine {
	router := getRouterEngine()

	// Using different paths for files
	router.LoadHTMLGlob("../static/html/**/*.html")

	// Loads the img, css and js folders
	router.Static("css/", "../static/css/")
	router.Static("js/", "../static/js/")
	router.Static("img/", "../static/img/")

	setRoutes(router)
	createSuperAdmin(config.DB)
	return router
}
