package handlers

import (
	"net/http"
	"net/url"
	"sigma/services/login"

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

func LoginPOST() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usern := ctx.Request.FormValue("nome_login")
		passwd := ctx.Request.FormValue("senha_cad")

		user := login.DefaultUserInfo()
		if !user.CheckLogin(usern, passwd) {
			ctx.HTML(
				http.StatusUnauthorized,
				"login.html",
				gin.H{
					"IsCorrect": "Usuário e/ou senha incorretos",
				},
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"msg": "Boa dia seu merda",
			},
		)

	}
}
