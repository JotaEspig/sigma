package handlers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func LoginRedirect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		location := url.URL{Path: "/login"}
		ctx.Redirect(http.StatusFound, location.RequestURI())
	}
}

func LoginGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK, "login.html", nil,
		)
	}
}
