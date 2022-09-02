package middlewares

import (
	"net/http"
	"sigma/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func abortWithHTML(ctx *gin.Context, status int, file string) {
	ctx.HTML(status, file, nil)
	ctx.Abort()
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		// checks if token is expired
		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		ctx.Set("username", claims["username"])
		ctx.Next()
	}
}

func IsStudentMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		if claims["type"] != "student" {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// checks if token is expired
		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		ctx.Set("username", claims["username"])
		ctx.Next()
	}
}

func IsTeacherMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		if claims["type"] != "teacher" {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// checks if token is expired
		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		ctx.Set("username", claims["username"])
		ctx.Next()
	}
}

func IsAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// check if token is valid
		dToken, err := config.JWTService.ValidateToken(token)
		if err != nil || !dToken.Valid {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)
		if err != nil || claims == nil {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		if claims["type"] != "admin" {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		// checks if token is expired
		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		ctx.Set("username", claims["username"])
		ctx.Next()
	}
}

// BE CAREFUL WITH THIS MIDDLEWARE,
// it needs to be used with the IsAdminMiddleware to work properly
func IsSuperAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.GetString("username")

		// the username being "admin" means that the user is a super admin
		if username != "admin" {
			abortWithHTML(ctx, http.StatusUnauthorized, "access_denied.html")
			return
		}

		ctx.Next()
	}
}
