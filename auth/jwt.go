package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Returns the secret key set in the environment
func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

// Values that will be contained in the token
type authClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}

// Parameters used in jwt authentication
type JWTService struct {
	SecretKey string
	Issuer    string
}

// Creates a default jwtService struct
func JWTAuthService() *JWTService {
	return &JWTService{
		SecretKey: getSecretKey(),
		Issuer:    "SIGMA",
	}
}

// Generates a token according to the username.
// Returns error if an error has occurred in getting the signed token
func (service *JWTService) GenerateToken(username string, isAdmin bool) (string, error) {
	claims := &authClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
		username,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedToken, err := token.SignedString([]byte(service.SecretKey))
	return encodedToken, err
}

// Validates the token, according to the secret key
func (service *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Checks if the token is valid trying to convert it to HMAC
		_, isValid := token.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(service.SecretKey), nil
	}

	return jwt.Parse(encodedToken, keyFunc)
}
