package controllers

import (
	"net/http"
	"sigma/db"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

// Validates an user
func ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")

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

// Gets public user info, according to request
func GetUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")

		resp := struct {
			Params []string
		}{}
		// ShouldBind is used to not set header status code to 400
		// if there is an error
		ctx.ShouldBindJSON(&resp)

		params := db.GetColumns(user.PublicUserParams, resp.Params...)
		u, err := user.GetUser(db.DB, username, params...)
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
