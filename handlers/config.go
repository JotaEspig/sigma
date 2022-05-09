package handlers

import (
	auth "sigma/services/authentication"
	"sigma/services/database"
)

// Connection and database variables
var Conn = database.ConnInit()
var db = Conn.GetDB()

// JWTService variable
var defaultJWT = auth.JWTAuthService()
