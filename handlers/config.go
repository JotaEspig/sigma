package handlers

import (
	auth "sigma/services/authentication"
	"sigma/services/database"
)

// Database variable
var db = database.ConnInit().GetDB()

// jwtService variable
var defaultJWT = auth.JWTAuthService()
