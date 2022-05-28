package controllers

import (
	"net/http"
	"sigma/config"
	"sigma/db"
	"sigma/models/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Validates an user
func ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		//dToken means decoded token
		dToken, err := config.DefaultJWT.ValidateToken(token)
		if err != nil || !dToken.Valid {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		u, err := user.GetUser(db.DB, claims["username"].(string))
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"claims": claims,
				"user":   u.ToMap(),
			},
		)
	}
}

// Gets public user info, according to request
func GetUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")

		resp := struct {
			Params []string
		}{}
		// ShouldBind is used to not set header status code to 400
		// if there is an error
		ctx.ShouldBindJSON(&resp)

		u, err := user.GetUser(db.DB, username, resp.Params...)
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"user": u.ToMap(),
			},
		)
	}
}
