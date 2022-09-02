package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/classroom"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddClassroom adds a classroom to the database
func AddClassroom() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := struct {
			Name string `json:"name"`
			Year int    `json:"year"`
		}{}
		err := ctx.BindJSON(&r)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c, err := classroom.InitClassroom(r.Name, uint16(r.Year))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = config.DB.Create(c).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// GetAllClassroomsInfo gets parcial information about every classroom
func GetAllClassroomsInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		classrooms := []classroom.Classroom{}
		err := config.DB.Find(&classrooms).Error
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

// GetClassroomInfo gets all information about a classroom
func GetClassroomInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		classroomID, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		c := classroom.Classroom{}
		err = config.DB.Preload("Students.User").First(&c, classroomID).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(http.StatusOK, c.ToMap())
	}
}

func init() {
	config.DB.AutoMigrate(&classroom.Classroom{})
}
