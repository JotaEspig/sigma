package studentauth

import (
	userauth "sigma/services/authentication/user"
	"sigma/services/database"
	"testing"
)

var db = database.ConnInit().GetDB()

// Constants to use in tests, def = default
const (
	defUsername = "defUsername"
	defPasswd   = "defPasswd"
	defName     = "defName"
	defSurname  = "defSurname"
	defEmail    = "defEmail"
)

func TestAddStudent(t *testing.T) {

	user := userauth.InitUser(defSurname, defEmail, defName, defSurname, defPasswd)
	stud := InitStudent(user)
	AddStudent(db, stud)

}
