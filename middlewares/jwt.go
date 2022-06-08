package middlewares

import (
	"net/http"
	"sigma/auth"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		token, err := ctx.Cookie("auth")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := auth.GetTokenClaims(token)
		if err != nil || claims == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims["username"] != username {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// checks if token is expired
		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}

func IsAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		token, err := ctx.Cookie("auth")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := auth.GetTokenClaims(token)
		if err != nil || claims == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims["username"] != username {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !claims["isAdmin"].(bool) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// checks if token is expired
		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
