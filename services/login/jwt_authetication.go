package login

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(string) string
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
	Username string `json:"name"`
	jwt.StandardClaims
}

// Creates a default jwtService struct
func JWTAuthService() *jwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "SIGMA",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

// Generates a token according to
func (service *jwtService) GenerateToken(username string) string {
	return ""
}

func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
}
