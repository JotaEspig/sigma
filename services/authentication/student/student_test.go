package studentauth

import (
	userauth "sigma/services/authentication/user"
	"sigma/services/database"
	"testing"
)

var db = database.ConnInit().GetDB()

// Constants to use in tests, def = default
const (
	defUsername   = "defUsername"
	defPasswd     = "defPasswd"
	defName       = "defName"
	defSurname    = "defSurname"
	defEmail      = "defEmail"
	nonExistentID = 9999
)

func TestAddStudent(t *testing.T) {
	u := userauth.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	userauth.AddUser(db, u)
	u, err := userauth.GetUser(db, u.Username)
	if err != nil {
		t.Error(err)
	}

	s, err := InitStudent(u)
	if err != nil {
		t.Error(err)
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("adding legit student: %s", r)
			}
		}()
		AddStudent(db, s)
	}()

	// repeating the same action
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("adding repeated student (it's not supposed to work)")
			}
		}()
		AddStudent(db, s)
	}()
}

func TestGetStudent(t *testing.T) {
	u := userauth.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		userauth.AddUser(db, u)
	}()

	u, err := userauth.GetUser(db, defUsername)
	if err != nil {
		t.Errorf("getting legit student: %s", err)
	}

	s, err := InitStudent(u)
	if err != nil {
		t.Error(err)
	}

	s.Year.Scan(2)
	s.Status.Scan("ativo")

	// Adds student
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Error(r)
			}
		}()
		AddStudent(db, s)
	}()

	_, err = GetStudent(db, u.ID)
	// Checks if get student parcial info is working
	if err != nil {
		t.Errorf("getting legit student (parcial info): %s", err)
	}

	s, err = GetStudent(db, u.ID, "year", "status")
	if err != nil {
		t.Errorf("getting legit student (parcial info): %s", err)
	}
	if s.Year.Int16 == 0 {
		t.Errorf("getting legit student (parcial info): year is empty")
	}
	if s.Status.String == "" {
		t.Errorf("getting legit student (parcial info): status is empty")
	}
	if s.ClassID.Int64 != 0 {
		t.Errorf("getting legit student (parcial info): class id is filled")
	}

	_, err = GetStudent(db, nonExistentID)
	if err == nil {
		t.Errorf("getting non existent student (it's not supposed to work): %s", err)
	}
}
