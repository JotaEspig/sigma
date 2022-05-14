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
	defEmail    = "defEmail"
	defName     = "defName"
)

func TestUserValidate(t *testing.T) {
	u1 := InitUser(defUsername, defEmail, defName, defPasswd)
	u2 := InitUser(defUsername, defEmail, defName, defPasswd)
	if !u1.Validate(u2.Username, defPasswd) {
		t.Error("validating user: 2 identical users couldn't pass the validation")
	}

	fakePasswd := "fake passwd"
	u3 := InitUser(defUsername, defEmail, defName, defPasswd)
	u4 := InitUser(defUsername, defEmail, defName, fakePasswd)
	if !u3.Validate(u4.Username, defPasswd) {
		t.Error("validating user: 2 different users could pass the validation (it's supposed to not)")
	}
}

func TestAddUser(t *testing.T) {
	u := InitUser(defUsername, defEmail, defName, defPasswd)

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
	_, err := GetUser(db, defUsername)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	_, err = GetUser(db, "non-existent-user")
	if err == nil {
		t.Errorf("getting non existent user (it's not supposed to work): %s", err)
	}
}

func TestRmUser(t *testing.T) {
	u := InitUser(defUsername, defEmail, defName, defPasswd)

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
