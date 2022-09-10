package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/admin"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddAdmin adds an admin to the database
func AddAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
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

		err = config.DB.Transaction(func(tx *gorm.DB) error {
			a.User.Type = "admin"
			err := config.DB.Model(a.User).Update("type", a.User.Type).Error
			if err != nil {
				return err
			}

			err = tx.Create(a).Error
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// GetAdminInfo gets an admin from the database using username store in the context
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

// UpdateAdmin updates an admin from the database using username store in the context
// (if not found, then from parameter at url)
func UpdateAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		newValues := admin.Admin{}
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

// DeleteAdmin deletes an admin from the database
func DeleteAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		err := admin.RmAdmin(config.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
