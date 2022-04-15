package login

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTDefault = JWTAuthService()

type JWTService interface {
	GenerateToken(string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
}

// Parameters used in jwt authentication
// 	secretKey : key used in the generation and validation of a token
// 	issuer : who issued the token
type jwtService struct {
	secretKey string
	issuer    string
}

// Values that will be contained in the token
type authClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Creates a default jwtService struct
func JWTAuthService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "SIGMA",
	}
}

// Returns the secret key set in the environment
func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

// Generates a token according to
func (service *jwtService) GenerateToken(username string) (string, error) {
	claims := &authClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	// Creates the token using HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Gets the string with the encoded token
	encodedToken, err := token.SignedString([]byte(service.secretKey))
	return encodedToken, err
}

// Validates the token, according to the secret key
func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	// Creates a key function
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Checks if the token is valid trying to convert it to HMAC
		_, isValid := token.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, fmt.Errorf("Invalid token")
		}
		// If it's valid, will return the secret key to the parser
		return []byte(service.secretKey), nil
	}

	return jwt.Parse(encodedToken, keyFunc)
}
