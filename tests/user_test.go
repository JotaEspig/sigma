package tests

import (
	"sigma/db"
	"sigma/models/user"
	"testing"
)

var db_ = db.ConnInit().GetDB()

// Constants to use in tests, def = default
const (
	defUsername = "defUsername"
	defPasswd   = "defPasswd"
	defName     = "defName"
	defSurname  = "defSurname"
	defEmail    = "defEmail"
)

func TestGetColumns(t *testing.T) {
	columns := []string{"username", "password"}

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Error(r)
			}
		}()

		newColumns := user.GetColumns(columns...).([]string)
		if newColumns[0] != "username" {
			t.Errorf("get columns: There is no username in first index")
		}
		if newColumns[1] != "password" {
			t.Errorf("get columns: There is no password in seconde index")
		}
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Error(r)
			}
		}()

		newColumns := user.GetColumns().(string)
		if newColumns != "*" {
			t.Errorf("get columns: It's not *")
		}
	}()

}

// TODO Jota: Must make tests use gorm instead of sqlx

func TestUserValidate(t *testing.T) {
	u1 := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	u2 := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	if !u1.Validate(u2.Username, defPasswd) {
		t.Error("validating user: 2 identical users couldn't pass the validation")
	}

	fakePasswd := "fake passwd"
	u3 := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	u4 := user.InitUser(defUsername, defEmail, defName, defSurname, fakePasswd)
	if !u3.Validate(u4.Username, defPasswd) {
		t.Error("validating user: 2 different users could pass the validation (it's supposed to not)")
	}
}

func TestAddUser(t *testing.T) {
	db_.AutoMigrate(&user.User{})
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("adding legit user: %s", r)
			}
		}()
		user.AddUser(db_, u)
	}()

	user.RmUser(db_, defUsername)
}

func TestGetUser(t *testing.T) {
	db_.AutoMigrate(&user.User{})
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		user.AddUser(db_, u)
	}()

	u = user.GetUser(db_, defUsername)
	if u.Model.ID == 0 {
		t.Errorf("getting legit user: ID is 0")
	}

	u = user.GetUser(db_, defUsername, "username", "email")
	// Checks if get user parcial info is working
	if u.Username == "" {
		t.Errorf("getting legit user (parcial info): username is empty")
	}
	if u.Email == "" {
		t.Errorf("getting legit user (parcial info): email is empty")
	}
	if u.Name != "" {
		t.Errorf("getting legit user (parcial info): name is filled")
	}

	u = user.GetUser(db_, "non-existent-user")
	if u.Model.ID != 0 {
		t.Errorf("getting non existent user (it's not supposed to work): ID is not 0")
	}

	user.RmUser(db_, defUsername)
}

func TestRmUser(t *testing.T) {
	db_.AutoMigrate(&user.User{})
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		user.AddUser(db_, u)
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("removing legit user: %s", r)
			}
		}()
		user.RmUser(db_, u.Username)
	}()
}
