package tests

import (
	"sigma/auth"
	"testing"
)

func TestGenerateAndValidate(t *testing.T) {
	// Correct way of using
	jwtServ := auth.JWTAuthService()
	eToken, err := jwtServ.GenerateToken("JotaEspig", "")
	if err != nil {
		t.Errorf("Error in generating token: %s", err)
	}

	_, err = jwtServ.ValidateToken(eToken)
	if err != nil {
		t.Errorf("Error in validating token: %s", err)
	}

	// testing with random digits
	_, err = jwtServ.ValidateToken("dad8i123131")
	if err == nil {
		t.Error("Error in validating a fake token")
	}

	// Testing with other secret key
	fakeJWTServ := &auth.JWTService{
		SecretKey: "secr",
		Issuer:    "Other",
	}
	eToken, err = fakeJWTServ.GenerateToken("JotaEspig", "")
	if err != nil {
		t.Errorf("Error in generating fake token (other secret key): %s", err)
	}
	dToken, err := jwtServ.ValidateToken(eToken) // Trying to validate token made with the fake
	if dToken.Valid || err == nil {
		t.Errorf("Error in validating fake token (other secret key): %s", err)
	}
}
