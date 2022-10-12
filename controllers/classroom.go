package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/classroom"
	"sigma/models/student"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddClassroom adds a classroom to the database. It receives "name" and "year" (uint16)
// as x-www-form-urlencoded
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

// GetClassroomInfo gets all information about a classroom. It receives "id" (int) from URL
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

// UpdateClassroom updates a classroom in the database. It receives "id" (int) from URL
// and "name" and "year" (uint16) as x-www-form-urlencoded
func UpdateClassroom() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		c := classroom.Classroom{}
		err := config.DB.Select("id").First(&c, id).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.Name = ctx.PostForm("name")
		yearStr := ctx.PostForm("year")
		if yearStr == "" {
			yearStr = "0" // To prevent BadRequest Error since user may send it empty
		}
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Year = uint16(year)
		err = config.DB.Updates(c).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// DeleteClassroom deletes a classroom from the database. It receives "id" (int) from URL
func DeleteClassroom() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		config.DB.Model(student.Student{}).Where("classroom_id = ?", id).
			Update("classroom_id", nil)

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
