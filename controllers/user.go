/*
The functions below are the functions that a user can call
*/
package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/admin"
	"sigma/models/student"
	"sigma/models/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Generic route for user, gets PUBLIC info of
// either user or its children (student, admin)
func GetPublicUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		u, err := user.GetUser(config.DB, username, "username", "type")

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		f := getPublicInfoFuncs[u.Type]
		f(ctx, u.Username)
	}
}

// Generic route for user, gets ALL info of
// either user or its children (student, admin)
func GetAllUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		u, err := user.GetUser(config.DB, username, "username", "type")
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		f := getAllInfoFuncs[u.Type]
		f(ctx, u.Username)
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
		dToken, err := config.JWTService.ValidateToken(token)
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

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"username": claims["username"],
				"type":     claims["type"],
			},
		)
	}
}

// Updates a user's info
func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newValues := user.User{}
		username := ctx.Param("username")
		u, err := user.GetUser(config.DB, username, "id")
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = ctx.ShouldBindJSON(&newValues)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		newValues.ID = u.ID
		// These lines exists to make sure that the user
		// is not changing the password, or type of the user
		if newValues.HashedPassword != "" {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if newValues.Type != "" {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = user.UpdateUser(config.DB, &newValues)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, nil)
	}
}

// Contains functions to get public info of
// either user or its children (student, admin)
// "" means user has no type
var getPublicInfoFuncs = map[string]func(*gin.Context, string){
	"": func(ctx *gin.Context, username string) {
		u, err := user.GetUser(config.DB, username,
			user.PublicUserParams...)

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
	},

	"student": func(ctx *gin.Context, username string) {
		s, err := student.GetStudent(config.DB, username,
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
	},

	"admin": func(ctx *gin.Context, username string) {
		a, err := admin.GetAdmin(config.DB, username,
			admin.PublicAdminParams...)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": a.ToMap(),
			},
		)
	},
}

// Contains functions to get all info of
// either user or its children (student, admin)
var getAllInfoFuncs = map[string]func(*gin.Context, string){
	"": func(ctx *gin.Context, username string) {
		u, err := user.GetUser(config.DB, username)

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
	},

	"student": func(ctx *gin.Context, username string) {
		s, err := student.GetStudent(config.DB, username)

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
	},

	"admin": func(ctx *gin.Context, username string) {
		a, err := admin.GetAdmin(config.DB, username)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": a.ToMap(),
			},
		)
	},
}
