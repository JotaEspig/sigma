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

		u, err := user.GetUser(config.DB, usern, "username", "password", "type")
		if err != nil || !u.Validate(usern, passwd) {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		token, err := config.JWTService.GenerateToken(u.Username, u.Type)
		if err != nil || token == "" {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"username": u.Username,
				"type":     u.Type,
				"token":    token,
			},
		)
	}
}
