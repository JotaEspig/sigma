package handlers

import "sigma/services/database"

// Database variable
var db = database.Conn.GetDB()
