package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/user"

	"github.com/gin-gonic/gin"
)

// Does the login process, it validates the user and password and return a token in JSON
func LoginPOST() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usern := ctx.PostForm("username")
		passwd := ctx.PostForm("password")

		user, err := user.GetUser(config.DB, usern)
		if err != nil || !user.Validate(usern, passwd) {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		token, err := config.JWTService.GenerateToken(usern, user.Type)
		if err != nil || token == "" {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"username": usern,
				"token":    token,
			},
		)
	}
}
