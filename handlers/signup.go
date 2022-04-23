package handlers

import (
	"net/http"
	"sigma/services/db"
	"sigma/services/login"

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
		passwd := ctx.PostForm("password")

		u := login.InitUser(usern, email, name, passwd)
		err := login.AddUser(db.DB, u)
		if err != nil {
			ctx.JSON(
				http.StatusConflict, nil,
			)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
