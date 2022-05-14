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

		user := auth.InitUser(usern, email, name, passwd)

		err := db.Ping() // Tests the database
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		// It will recover if an error occurs in AddUser
		// that means that duplicate key error happened
		defer func() {
			if r := recover(); r != nil {
				ctx.Status(http.StatusConflict)
			}
		}()
		auth.AddUser(db, user)

		ctx.Status(http.StatusOK)
	}
}
