/*
The functions below are the functions that an admin can call
*/

package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/admin"

	"github.com/gin-gonic/gin"
)

func GetAdminInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		a, err := admin.GetAdmin(config.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(http.StatusOK, a.ToMap())
	}
}

func UpdateAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newValues := admin.Admin{}
		username := ctx.Param("username")
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

func DeleteAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		err := admin.RmAdmin(config.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
