package handlers

import (
	"net/http"
	"net/url"
	"sigma/services/login"

	"github.com/gin-gonic/gin"
)

// Just redirect the user to the login page
func LoginRedirect(ctx *gin.Context) {
	location := url.URL{Path: "/login"}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}

// At the moment, this function just serves the html file
func LoginGET(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK, "login.html", nil,
	)
}

func LoginPOST(ctx *gin.Context) {
	usern := ctx.PostForm("nome_login")
	passwd := ctx.PostForm("senha_cad")

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
