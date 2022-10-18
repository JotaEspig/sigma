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

type userImpl interface {
	ToMap() map[string]interface{}
}

// GetPublicUserInfo is a controller that gets PUBLIC info of
// either user or its children (student, admin)
func GetPublicUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		err := config.DB.Select("id", "type").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		// Gets the specific function according to the user type
		f := getPublicInfoFuncs[u.Type]
		uImpl, err := f(u.ID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": uImpl.ToMap(),
			},
		)
	}
}

// GetAllUserInfo is a controller that gets ALL info of
// either user or its children (student, admin)
func GetAllUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		err := config.DB.Select("id", "type").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		// Gets the specific function according to the user type
		f := getAllInfoFuncs[u.Type]
		uImpl, err := f(u.ID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": uImpl.ToMap(),
			},
		)
	}
}

// SearchUsers is a controller that searchs for user using ILIKE clause
func SearchUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users := []user.User{}
		username := getUsername(ctx)
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

// UpdateUser is a controller that updates a logged user's info WITH RESTRICTIONS
func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := getUsername(ctx)
		u := user.User{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		u.Email = ctx.PostForm("email")
		u.Name = ctx.PostForm("name")
		u.Surname = ctx.PostForm("surname")

		err = config.DB.Updates(&u).Error
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
var getPublicInfoFuncs = map[string]func(uint) (userImpl, error){
	"": func(id uint) (userImpl, error) {
		u := user.User{}
		err := config.DB.Where("id = ?", id).First(&u).Error
		return u, err
	},

	"student": func(id uint) (userImpl, error) {
		s := student.Student{}
		err := config.DB.Preload("User").Where("id = ?", id).First(&s).Error
		return s, err
	},

	"teacher": func(id uint) (userImpl, error) {
		t := teacher.Teacher{}
		err := config.DB.Preload("User").Where("id = ?", id).First(&t).Error
		return t, err
	},

	"admin": func(id uint) (userImpl, error) {
		a := admin.Admin{}
		err := config.DB.Preload("User").Where("id = ?", id).First(&a).Error
		return a, err
	},
}

// Contains functions to get all info of
// either user or its children (student, admin)
var getAllInfoFuncs = map[string]func(uint) (userImpl, error){
	"": func(id uint) (userImpl, error) {
		u := user.User{}
		err := config.DB.Where("id = ?", id).First(&u).Error
		return u, err
	},

	"student": func(id uint) (userImpl, error) {
		s := student.Student{}
		err := config.DB.Preload("User").Where("id = ?", id).First(&s).Error
		return s, err
	},

	"teacher": func(id uint) (userImpl, error) {
		t := teacher.Teacher{}
		err := config.DB.Preload("User").Where("id = ?", id).First(&t).Error
		return t, err
	},

	"admin": func(id uint) (userImpl, error) {
		a := admin.Admin{}
		err := config.DB.Preload("User").Where("id = ?", id).First(&a).Error
		return a, err
	},
}

// AutoMigrate the user table
func init() {
	config.DB.AutoMigrate(&user.User{})
}
