package auth

import (
	"sigma/services/database"
	"testing"
)

var db = database.ConnInit().GetDB()

func TestUserValidate(t *testing.T) {
	passwd := "admin"

	u1 := InitUser("admin", "admin@admin.com", "admin", passwd)
	u2 := InitUser("admin", "admin@admin.com", "admin", passwd)
	if !u1.Validate(u2.Username, passwd) {
		t.Error("validating user: 2 identical users couldn't pass the validation")
	}

	fakePasswd := "fake passwd"
	u3 := InitUser("admin", "admin@admin.com", "admin", passwd)
	u4 := InitUser("admin", "admin@admin.com", "admin", fakePasswd)
	if !u3.Validate(u4.Username, passwd) {
		t.Error("validating user: 2 different users could pass the validation (it's supposed to not)")
	}
}

func TestAddUser(t *testing.T) {
	u := InitUser("admin", "admin@admin.com", "admin", "admin")

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
				t.Errorf(
					"adding repeated user (it's not supposed to work): %s",
					r,
				)
			}
		}()
		AddUser(db, u)
	}()
}

func TestGetUser(t *testing.T) {
	_, err := GetUser(db, "admin")
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	_, err = GetUser(db, "non-existent-user")
	if err == nil {
		t.Errorf("getting non existent user (it's not supposed to work): %s", err)
	}
}
