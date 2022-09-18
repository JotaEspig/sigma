package controllers

import "github.com/gin-gonic/gin"

// TODO Check this function to be compatible with target url param
// maybe change "target" to "username" in url params

func getUsername(ctx *gin.Context) string {
	username := ctx.Param("username")
	if username == "" {
		username = ctx.GetString("username")
	}

	return username
}
