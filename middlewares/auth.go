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

// AuthMiddleware is a middleware to check if a user is logged or not
// and it set the username to the context to be used in next processes
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

// IsStudentMiddleware is a middleware to check if the user is logged and if it's a student or not
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

// IsTeacherMiddleware is a middleware to check if the user is logged and if it's a teacher or not
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

// IsAdminMiddleware is a middleware to check if the user is logged and if it's a admin or not
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

// IsSuperAdminMiddleware is a middleware that adds one more layer to check
// if username equals "admin".
// It needs to be used with the IsAdminMiddleware to work properly
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
