package handlers

import (
	jwtauth "sigma/services/authentication/jwt"
	"sigma/services/database"
)

// Database variables
var db = database.ConnInit().GetDB()

// JWTService variable
var defaultJWT = jwtauth.JWTAuthService()
