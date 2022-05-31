package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/db"
	"sigma/models/student"
	"sigma/models/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Serves "aluno.html" page
func GetAlunoPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "aluno.html", nil)
	}
}

// Serves "alunoinfo.html" page
func GetUserInfoPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "alunoinfo.html", nil)
	}
}

// Gets public user info, according to request
func GetPublicUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		u, err := user.GetUser(db.DB, username, user.PublicUserParams...)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		if u.Type == "" {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"user": u.ToMap(),
				},
			)
			return
		}

		switch u.Type {
		case "student":
			s, err := student.GetStudent(db.DB, username,
				student.PublicStudentParams...)

			if err != nil {
				ctx.AbortWithStatus(http.StatusNotFound)
				return
			}

			ctx.JSON(
				http.StatusOK,
				gin.H{
					"user": s.ToMap(),
				},
			)
			return
		}
	}
}

// Gets all user info, according to request
func GetAllUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")

		u, err := user.GetUser(db.DB, username, user.UserParams...)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		if u.Type == "" {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"user": u.ToMap(),
				},
			)
			return
		}

		switch u.Type {
		case "student":
			s, err := student.GetStudent(db.DB, username,
				student.StudentParams...)
			if err != nil {
				ctx.AbortWithStatus(http.StatusNotFound)
				return
			}

			ctx.JSON(
				http.StatusOK,
				gin.H{
					"user": s.ToMap(),
				},
			)
			return
		}
	}
}

// Validates a user with token got from cookie auth
func ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// check if token is valid
		dToken, err := config.DefaultJWT.ValidateToken(token)
		if err != nil || !dToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		username := claims["username"].(string)

		u, err := user.GetUser(db.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": u.ToMap(),
			},
		)
	}
}
