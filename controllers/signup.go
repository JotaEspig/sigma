package controllers

import (
	"fmt"
	"net/http"
	"sigma/config"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

// SignupPOST gets the values from the form, creates a user and inserts it in the database
func SignupPOST() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usern := ctx.PostForm("username")
		email := ctx.PostForm("email")
		name := ctx.PostForm("name")
		surname := ctx.PostForm("surname")
		passwd := ctx.PostForm("password")

		u := user.InitUser(usern, email, name, surname, passwd)
		fmt.Println(u)

		err := config.DB.Create(u).Error
		if err != nil {
			ctx.Status(http.StatusConflict)
			fmt.Println(err)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
