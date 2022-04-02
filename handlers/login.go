package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(
			http.StatusOK, "login.html", nil,
		)
	}
}
