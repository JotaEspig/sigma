package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Serves "login" page
func GetLoginPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	}
}

// Serves "cadastro.html" page
func SignupGET() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "cadastro.html", nil)
	}
}

// Serves "profile.html" page
func GetProfilePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "profile.html", nil)
	}
}

// Serves "user.html" page
func GetUserPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "user.html", nil)
	}
}

// Serves "aluno.html" page
func GetStudentPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "aluno.html", nil)
	}
}

// Serves "teacher.html" page
func GetTeacherPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "teacher.html", nil)
	}
}

// Server "admin.html" page
func GetAdminPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin.html", nil)
	}
}
