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

// Serves "alunoinfo.html" page
func GetUserInfoPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "alunoinfo.html", nil)
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

		params := db.GetColumns(user.PublicUserParams, resp.Params...)
		u, err := user.GetUser(db.DB, username, params...)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
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

// Serves "aluno.html" page
func GetAlunoPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "aluno.html", nil)
	}
}

// Validates an user
func ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")

		u, err := user.GetUser(db.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
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

// Validates an user with token got from cookie auth
func ValidateUserWithToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// check if token is valid
		dToken, err := config.DefaultJWT.ValidateToken(token)
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

		username := claims["username"].(string)

		u, err := user.GetUser(db.DB, username)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
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
