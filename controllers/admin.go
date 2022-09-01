package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/admin"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

// Adds an admin to the database
func AddAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		u := user.User{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		a, err := admin.InitAdmin(&u)
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

// Gets an admin from the database using username store in the context
// (if not found, then from parameter at url)
func GetAdminInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.GetString("username")
		if username == "" {
			username = ctx.Param("username")
		}

		a, err := admin.GetAdmin(config.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(http.StatusOK, a.ToMap())
	}
}

// Updates an admin from the database using username store in the context
// (if not found, then from parameter at url)
func UpdateAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newValues := admin.Admin{}
		username := ctx.GetString("username")
		if username == "" {
			username = ctx.Param("username")
		}

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

// Deletes an admin from the database
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
