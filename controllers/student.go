package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/student"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

// GetStudentInfo gets student info according to the username
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

		ctx.JSON(http.StatusOK, s.ToMap())
	}
}

// AutoMigrate the student table
func init() {
	config.DB.AutoMigrate(&student.Student{})
}
