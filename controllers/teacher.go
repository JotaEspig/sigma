package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/teacher"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddTeacher is a controller that changes the user type to "teacher"
// and adds a teacher row to the database if it doesn't already exists
func AddTeacher() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = config.DB.Transaction(func(tx *gorm.DB) error {
			u.Type = "teacher"
			err := tx.Updates(&u).Error
			if err != nil {
				return err
			}

			// Checks if a teacher already exists with this ID.
			// If it exists, it doesn't create another one
			t := &teacher.Teacher{}
			tx.First(t, u.ID)
			if t.UID != 0 {
				return nil
			}

			t, err = teacher.InitTeacher(&u)
			if err != nil {
				return err
			}

			return tx.Create(t).Error
		})
	}
}

// GetTeacherInfo is a controller that gets teacher info according to the username
func GetTeacherInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		t := teacher.Teacher{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = config.DB.Preload("User").Where("id = ?", u.ID).First(&t).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"teacher": t.ToMap(),
			},
		)
	}
}

// UpdateTeacher is a controller that updates the teacher that is logged in
func UpdateTeacher() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		t := teacher.Teacher{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = config.DB.Select("id").First(&t, u.ID).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		t.Education = ctx.PostForm("education")
		err = config.DB.Omit("id").Updates(t).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// DeleteTeacher is a controller that deletes a teacher from the database
func DeleteTeacher() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = config.DB.Transaction(func(tx *gorm.DB) error {
			err = config.DB.Model(&u).Update("type", "").Error
			if err != nil {
				return err
			}

			return config.DB.Unscoped().Delete(&teacher.Teacher{}, u.ID).Error
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func init() {
	config.DB.AutoMigrate(&teacher.Teacher{})
}
