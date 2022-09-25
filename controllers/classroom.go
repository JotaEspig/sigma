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
		name := ctx.PostForm("name")
		yearStr := ctx.PostForm("year")
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c, err := classroom.InitClassroom(name, uint16(year))
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

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"classrooms": classroomsMap,
			},
		)
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

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"classroom": c.ToMap(),
			},
		)
	}
}

// UpdateClassroom updates the values of name and year of a classroom
func UpdateClassroom() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		c := classroom.Classroom{}

		err := config.DB.First(&c, id).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.Name = ctx.PostForm("name")
		year, err := strconv.Atoi(ctx.PostForm("year"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadGateway)
			return
		}

		c.Year = uint16(year)
		err = config.DB.Updates(&c).Error
	}
}

// DeleteClassroom deletes a classroom from the database
func DeleteClassroom() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := config.DB.Unscoped().Delete(&classroom.Classroom{}, id).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
		}
		ctx.Status(http.StatusOK)
	}
}

func init() {
	config.DB.AutoMigrate(&classroom.Classroom{})
}
