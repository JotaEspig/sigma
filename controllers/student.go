package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/student"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddStudent is a controller that changes the user type to "student"
// and adds a student row to the database if it doesn't already exists
func AddStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = config.DB.Transaction(func(tx *gorm.DB) error {
			u.Type = "student"
			err := tx.Updates(&u).Error
			if err != nil {
				return err
			}

			// Checks if a student already exists with this ID.
			// If it exists, it doesn't create another one
			s := &student.Student{}
			tx.First(s, u.ID)
			if s.UID != 0 {
				return nil
			}

			s, err = student.InitStudent(&u)
			if err != nil {
				return err
			}

			return tx.Create(s).Error
		})
	}
}

// GetStudentInfo is a controller that gets student info according to the username
func GetStudentInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		s := student.Student{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = config.DB.Preload("User").Where("id = ?", u.ID).First(&s).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"student": s.ToMap(),
			},
		)
	}
}

// UpdateStudent is a controller that updates the student status
func UpdateStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		s := student.Student{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = config.DB.Select("id").Where("id = ?", u.ID).First(&s).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		s.Status = ctx.PostForm("status")
		err = config.DB.Model(s).Omit("id").Updates(s).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// DeleteStudent is a controller that deletes a student from the database
func DeleteStudent() gin.HandlerFunc {
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

			return config.DB.Unscoped().Delete(&student.Student{}, u.ID).Error
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// AutoMigrate the student table
func init() {
	config.DB.AutoMigrate(&student.Student{})
}
