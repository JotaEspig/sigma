package teacher

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
)

// Slice of all teacher params
var TeacherParams = []string{
	"id",
	"education",
}

// Slice of public teacher params
var PublicTeacherParams = []string{
	"id",
	"education",
}

func init() {
	config.DB.AutoMigrate(&Teacher{})
}
