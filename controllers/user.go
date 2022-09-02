package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/admin"
	"sigma/models/student"
	"sigma/models/teacher"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

// GetPublicUserInfo is a generic route for user, gets PUBLIC info of
// either user or its children (student, admin)
func GetPublicUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		u := user.User{}
		err := config.DB.Select("username", "type").Where("username = ?", username).First(&u).Error

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		// Gets the specific function according to the user type
		f := getPublicInfoFuncs[u.Type]
		f(ctx, u.Username)
	}
}

// GetAllUserInfo is a generic route for user, gets ALL info of
// either user or its children (student, admin)
func GetAllUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.GetString("username")
		u := user.User{}
		err := config.DB.Select("username", "type").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		// Gets the specific function according to the user type
		f := getAllInfoFuncs[u.Type]
		f(ctx, u.Username)
	}
}

// SearchUsers searchs for user using ILIKE clause
func SearchUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users := []user.User{}
		username := ctx.Param("username")
		err := config.DB.Select("id", "username").Where("username ILIKE ?", "%"+username+"%").
			Limit(10).Find(&users).Error
		if err != nil || len(users) == 0 {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		usersMap := make([]map[string]interface{}, len(users))
		for i, u := range users {
			usersMap[i] = u.ToMap()
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"users": usersMap,
			},
		)
	}
}

// UpdateUser updates a logged user's info WITH RESTRICTIONS
func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newValues := user.User{}
		username := ctx.GetString("username")
		u := user.User{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		err = ctx.ShouldBindJSON(&newValues)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		newValues.ID = u.ID
		// These lines exists to make sure that the user
		// is not changing the password, or type of the user
		newValues.HashedPassword = ""
		newValues.Type = ""

		err = user.UpdateUser(config.DB, &newValues)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

// <>----<>----<>----<>----<>----<>----<>----<>----<>----<>----<>----<>----<>----<>
// The maps below are used to get the correct function according to the user type
// because the functions need to act differently depending on the user type
// <>----<>----<>----<>----<>----<>----<>----<>----<>----<>----<>----<>----<>----<>

// Contains functions to get public info of
// either user or its children (student, admin)
// "" means user has no type
var getPublicInfoFuncs = map[string]func(*gin.Context, string){
	"": func(ctx *gin.Context, username string) {
		u := user.User{}
		err := config.DB.Omit("password").Where("username = ?", username).First(&u).Error

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": u.ToMap(),
			},
		)
	},

	"student": func(ctx *gin.Context, username string) {
		s, err := student.GetStudent(config.DB, username,
			student.PublicStudentParams...)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": s.ToMap(),
			},
		)
	},

	"teacher": func(ctx *gin.Context, username string) {
		t, err := teacher.GetTeacher(config.DB, username,
			teacher.PublicTeacherParams...)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": t.ToMap(),
			},
		)
	},

	"admin": func(ctx *gin.Context, username string) {
		a, err := admin.GetAdmin(config.DB, username,
			admin.PublicAdminParams...)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": a.ToMap(),
			},
		)
	},
}

// Contains functions to get all info of
// either user or its children (student, admin)
var getAllInfoFuncs = map[string]func(*gin.Context, string){
	"": func(ctx *gin.Context, username string) {
		u := user.User{}
		err := config.DB.Where("username = ?", username).First(&u).Error

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": u.ToMap(),
			},
		)
	},

	"student": func(ctx *gin.Context, username string) {
		s, err := student.GetStudent(config.DB, username)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": s.ToMap(),
			},
		)
	},

	"teacher": func(ctx *gin.Context, username string) {
		t, err := teacher.GetTeacher(config.DB, username)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": t.ToMap(),
			},
		)
	},

	"admin": func(ctx *gin.Context, username string) {
		a, err := admin.GetAdmin(config.DB, username)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": a.ToMap(),
			},
		)
	},
}
