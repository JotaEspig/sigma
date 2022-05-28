package controllers

import (
	"net/http"
	"sigma/db"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

// At the moment, this function just serves the html file
func SignupGET() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK, "cadastro.html", nil,
		)
	}
}

// Gets the values from the form, creates a user and inserts it in the database
func SignupPOST() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usern := ctx.PostForm("username")
		email := ctx.PostForm("email")
		name := ctx.PostForm("name")
		surname := ctx.PostForm("surname")
		passwd := ctx.PostForm("password")

		u := user.InitUser(usern, email, name, surname, passwd)

		err := user.AddUser(db.DB, u)
		if err != nil {
			ctx.Status(http.StatusConflict)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
