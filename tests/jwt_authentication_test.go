package tests

import (
	"sigma/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidate(t *testing.T) {
	// Correct way of using
	jwtServ := auth.JWTAuthService()
	eToken, err := jwtServ.GenerateToken("JotaEspig", "")
	assert.Equal(t, nil, err)

	_, err = jwtServ.ValidateToken(eToken)
	assert.Equal(t, nil, err)

	// testing with random digits
	_, err = jwtServ.ValidateToken("dad8i123131")
	assert.NotEqual(t, nil, err)

	// Testing with other secret key
	fakeJWTServ := &auth.JWTService{
		SecretKey: "secr",
		Issuer:    "Other",
	}
	eToken, err = fakeJWTServ.GenerateToken("JotaEspig", "")
	assert.Equal(t, nil, err)
	dToken, err := jwtServ.ValidateToken(eToken) // Trying to validate token made with the fake
	assert.NotEqual(t, true, dToken.Valid)
	assert.NotEqual(t, nil, err)
}
