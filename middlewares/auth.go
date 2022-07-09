package middlewares

import (
	"net/http"
	"sigma/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AbortWithHTML(ctx *gin.Context, status int, file string) {
	ctx.HTML(status, file, nil)
	ctx.Abort()
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		// checks if token is expired
		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		ctx.Set("username", claims["username"])
		ctx.Next()
	}
}

func IsStudentMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		if claims["username"] != username {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		if claims["type"] != "student" {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// checks if token is expired
		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
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
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)
		if err != nil || claims == nil {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		if claims["username"] != username {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		if claims["type"] != "admin" {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// checks if token is expired
		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		ctx.Next()
	}
}

// BE CAREFUL WITH THIS MIDDLEWARE,
// it need to be used with the IsAdminMiddleware to work properly
func IsSuperAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		if username != "SUPERADMIN" {
			AbortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		ctx.Next()
	}
}
