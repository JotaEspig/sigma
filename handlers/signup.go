package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignupGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK, "cadastro.html", nil,
		)
	}
}
