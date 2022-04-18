package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// At the moment, this function just serves the html file
func SignupGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK, "cadastro.html", nil,
		)
	}
}
