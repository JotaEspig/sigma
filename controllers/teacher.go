package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/teacher"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

/*
return db.Transaction(func(tx *gorm.DB) error {
		t.User.Type = "teacher"
		err := db.Model(t.User).Omit("username", "password", "type").Updates(t.User).Error
		if err != nil {
			return err
		}

		err = tx.Create(t).Error
		if err != nil {
			return err
		}

		return nil
	})
*/

// GetTeacherInfo gets teacher info according to the username
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

// UpdateTeacher updates the teacher that is logged in
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

func init() {
	config.DB.AutoMigrate(&teacher.Teacher{})
}
