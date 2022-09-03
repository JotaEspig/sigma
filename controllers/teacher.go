package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/teacher"

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
		username := ctx.GetString("username")
		t, err := teacher.GetTeacher(config.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(http.StatusOK, t.ToMap())
	}
}

// UpdateTeacher updates the teacher that is logged in
func UpdateTeacher() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newValues := teacher.Teacher{}
		username := ctx.GetString("username")
		t, err := teacher.GetTeacher(config.DB, username, "id")
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.ShouldBindJSON(&newValues)
		newValues.UID = t.UID
		err = teacher.UpdateTeacher(config.DB, &newValues)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
