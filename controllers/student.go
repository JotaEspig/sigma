package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/student"

	"github.com/gin-gonic/gin"
)

func GetStudentInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.GetString("username")
		s, err := student.GetStudent(config.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(http.StatusOK, s.ToMap())
	}
}
