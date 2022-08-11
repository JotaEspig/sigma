package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/admin"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

// Adds an admin to the database using target param
func AddTargetAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("target")
		u, err := user.GetUser(config.DB, username, "id")
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		a, err := admin.InitAdmin(u)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = admin.AddAdmin(config.DB, a)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.Status(http.StatusOK)
	}
}

// Updates an admin from the database using target param
func UpdateTargetAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newValues := admin.Admin{}
		username := ctx.Param("target")
		a, err := admin.GetAdmin(config.DB, username, "id")
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.ShouldBindJSON(&newValues)
		newValues.UID = a.UID
		err = admin.UpdateAdmin(config.DB, &newValues)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// Deletes an admin from the database using target param
func DeleteTargetAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("target")
		err := admin.RmAdmin(config.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
