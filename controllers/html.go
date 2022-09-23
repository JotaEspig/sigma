package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServePage serves a html file
func ServePage(filename string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, filename, nil)
	}
}
