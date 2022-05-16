package handlers

import (
	auth "sigma/services/authentication/jwt"
	"sigma/services/database"
)

// Database variables
var db = database.ConnInit().GetDB()

// JWTService variable
var defaultJWT = auth.JWTAuthService()
