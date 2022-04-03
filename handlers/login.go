package handlers

import (
	"net/http"
	"net/url"

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
func LoginGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK, "login.html", nil,
		)
	}
}
