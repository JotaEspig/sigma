package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/classroom"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddClassroom() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		year, err := strconv.ParseUint(ctx.PostForm("year"), 10, 32)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c, err := classroom.InitClassroom(name, uint16(year))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = classroom.AddClassroom(config.DB, c)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func GetAllClassroomsInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		classrooms, err := classroom.GetAllClassrooms(config.DB)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		classroomsMap := make([]map[string]interface{}, len(classrooms))
		for i, c := range classrooms {
			classroomsMap[i] = c.ToMap()
		}

		ctx.JSON(http.StatusOK, classroomsMap)
	}
}

func GetClassroomInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		classroom_id, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		c, err := classroom.GetClassroom(config.DB, uint(classroom_id))
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(http.StatusOK, c.ToMap())
	}
}
