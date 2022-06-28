package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/admin"

	"github.com/gin-gonic/gin"
)

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
