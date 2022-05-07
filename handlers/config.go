package handlers

import (
	auth "sigma/services/authentication"
	"sigma/services/database"
)

// Database variable
var Conn = database.ConnInit()
var db = Conn.GetDB()

// jwtService variable
var defaultJWT = auth.JWTAuthService()
