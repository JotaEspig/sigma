package auth

import (
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

func TestUserValidate(t *testing.T) {
	u1 := InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	u2 := InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	if !u1.Validate(u2.Username, defPasswd) {
		t.Error("validating user: 2 identical users couldn't pass the validation")
	}

	fakePasswd := "fake passwd"
	u3 := InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	u4 := InitUser(defUsername, defEmail, defName, defSurname, fakePasswd)
	if !u3.Validate(u4.Username, defPasswd) {
		t.Error("validating user: 2 different users could pass the validation (it's supposed to not)")
	}
}

func TestAddUser(t *testing.T) {
	u := InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("adding legit user: %s", r)
			}
		}()
		AddUser(db, u)
	}()

	// repeating the same action
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("adding repeated user (it's not supposed to work)")
			}
		}()
		AddUser(db, u)
	}()
}

func TestGetUser(t *testing.T) {
	TestAddUser(t)

	_, err := GetUser(db, defUsername)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	u, err := GetUser(db, defUsername, "username", "email")
	// Checks if get user parcial info is working
	if err != nil {
		t.Errorf("getting legit user (parcial info): %s", err)
	}
	if u.Username == "" {
		t.Errorf("getting legit user (parcial info): username is empty")
	}
	if u.Email == "" {
		t.Errorf("getting legit user (parcial info): email is empty")
	}
	if u.Name != "" {
		t.Errorf("getting legit user (parcial info): name is filled")
	}

	_, err = GetUser(db, "non-existent-user")
	if err == nil {
		t.Errorf("getting non existent user (it's not supposed to work): %s", err)
	}
}

func TestRmUser(t *testing.T) {
	u := InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		AddUser(db, u)
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("removing legit user: %s", r)
			}
		}()
		RmUser(db, u.Username)
	}()
}
