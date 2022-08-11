package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/models/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Does the login process, it validates the user and password and return a token in JSON
func Login() gin.HandlerFunc {
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
				"type":  u.Type,
				"token": token,
			},
		)
	}
}

// If user it's logged, it sends JSON with username and type of the user
func IsLogged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"username": claims["username"],
				"type":     claims["type"],
			},
		)
	}
}
