package handlers

import (
	"net/http"
	auth "sigma/services/authentication"

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

		u := auth.InitUser(usern, email, name, passwd)
		err := auth.AddUser(db, u)
		if err != nil {
			ctx.Status(http.StatusConflict)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
