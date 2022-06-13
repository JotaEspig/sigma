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
