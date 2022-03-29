package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"Texto": gin.H{
					"Min": "Teste2",
				},
			},
		)
	}
}
