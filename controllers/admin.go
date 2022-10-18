package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/admin"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddAdmin is a controller that adds an admin to the database
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

// GetAdminInfo is a controller that gets an admin from the database using username
// stored in the context (if not found, then from parameter at url)
func GetAdminInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		a := admin.Admin{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = config.DB.Preload("User").First(&a, u.ID).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(http.StatusOK, a.ToMap())
	}
}

// UpdateAdmin is a controller that updates an admin from the database using username
// stored in the context (if not found, then from parameter at url)
func UpdateAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		a := admin.Admin{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = config.DB.Select("id").First(&a, u.ID).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		a.Role = ctx.PostForm("role")
		err = config.DB.Updates(a).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// DeleteAdmin is a controller that deletes an admin from the database
func DeleteAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		err := config.DB.Transaction(func(tx *gorm.DB) error {
			u := user.User{}
			err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
			if err != nil {
				return err
			}

			// Updates the type of the user to be empty
			err = config.DB.Model(&u).Update("type", "").Error
			if err != nil {
				return err
			}

			return config.DB.Unscoped().Delete(&admin.Admin{}, u.ID).Error
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
