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

// Gets the values from the form, creates an user and inserts it in the database
func SignupPOST() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usern := ctx.PostForm("username")
		email := ctx.PostForm("email")
		name := ctx.PostForm("name")
		surname := ctx.PostForm("surname")
		passwd := ctx.PostForm("password")

		u := user.InitUser(usern, email, name, surname, passwd)

		// It will recover if an error occurs in AddUser
		// that means that duplicate key error happened
		defer func() {
			if r := recover(); r != nil {
				ctx.Status(http.StatusConflict)
			}
		}()
		user.AddUser(db.DB, u)

		ctx.Status(http.StatusOK)
	}
}
