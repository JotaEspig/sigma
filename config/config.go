package config

import (
	"sigma/auth"
	"sigma/db"
)

// JWTService variable
var JWTService = auth.JWTAuthService()

// DB variable
var DB = db.ConnInit().GetDB()
