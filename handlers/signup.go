package handlers

import (
	"errors"
	"net/http"
	auth "sigma/services/authentication"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

const (
	codeDuplicateKey = "23505"
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
		err := auth.AddUser(db, user)
		if err != nil {
			var pqError *pq.Error
			errors.As(err, &pqError)
			if pqError.Code == codeDuplicateKey {
				ctx.Status(http.StatusConflict)
				return
			}
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
