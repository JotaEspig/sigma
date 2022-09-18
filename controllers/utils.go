package controllers

import "github.com/gin-gonic/gin"

func getUsername(ctx *gin.Context) string {
	username := ctx.Param("username")
	if username == "" {
		username = ctx.GetString("username")
	}

	return username
}
