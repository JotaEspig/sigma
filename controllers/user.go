/*
The functions below are the functions that a user can call
*/
package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/admin"
	"sigma/models/student"
	"sigma/models/teacher"
	"sigma/models/user"

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
		username := ctx.GetString("username")
		u, err := user.GetUser(config.DB, username, "username", "type")
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		f := getAllInfoFuncs[u.Type]
		f(ctx, u.Username)
	}
}

// Updates a user's info WITH RESTRICTIONS
func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newValues := user.User{}
		username := ctx.GetString("username")
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

		ctx.Status(http.StatusOK)
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

	"teacher": func(ctx *gin.Context, username string) {
		t, err := teacher.GetTeacher(config.DB, username,
			teacher.PublicTeacherParams...)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": t.ToMap(),
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

	"teacher": func(ctx *gin.Context, username string) {
		t, err := teacher.GetTeacher(config.DB, username)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": t.ToMap(),
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
