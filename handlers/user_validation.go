package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Gets the token and sends a JSON containing information about the user to the browser
// if the token is valid
func ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Defines a function that sends an unauthorized code to the server
		unauthorizedJSON := func(json gin.H) {
			ctx.JSON(
				http.StatusUnauthorized,
				json,
			)
		}

		response := struct {
			Token string `json:"token"`
		}{}

		ctx.BindJSON(&response)
		if response.Token == "" {
			unauthorizedJSON(nil)
			return
		}

		//dToken means decoded token
		dToken, err := defaultJWT.ValidateToken(response.Token)
		if err != nil || !dToken.Valid {
			unauthorizedJSON(nil)
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if err != nil {
			unauthorizedJSON(nil)
			return
		}
		if float64(now) > expiresAt {
			unauthorizedJSON(nil)
			return
		}

		content := make(gin.H)
		for key, value := range claims {
			content[key] = value
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"username": claims["username"],
			},
		)
	}
}