package controllers

import (
	"net/http"
	"net/url"
	"sigma/config"
	"sigma/db"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

// Just redirect the user to the login page
func LoginRedirect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		location := url.URL{Path: "/login"}
		ctx.Redirect(http.StatusFound, location.RequestURI())
	}
}

// At the moment, this function just serves the html file
func LoginGET() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK, "login.html", nil,
		)
	}
}

// Does the login process, it validates the user and password and return a token in JSON
func LoginPOST() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usern := ctx.PostForm("username")
		passwd := ctx.PostForm("password")

		user := user.GetUser(db.DB, usern)
		if !user.Validate(usern, passwd) {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		token, err := config.DefaultJWT.GenerateToken(usern)
		if err != nil || token == "" {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"token": token,
			},
		)
	}
}
