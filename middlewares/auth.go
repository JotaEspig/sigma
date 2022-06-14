package middlewares

import (
	"net/http"
	"sigma/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

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

func IsStudentMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		if claims["username"] != username {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims["type"] != "student" {
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
		if token == "" || err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)
		if err != nil || claims == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims["username"] != username {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims["type"] != "admin" {
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
