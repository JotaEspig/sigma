package tests

import (
	"sigma/models/user"
	"testing"
)

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

func TestGetUser(t *testing.T) {
	t.Skip("not implemented yet")
}

func TestUpdateUser(t *testing.T) {
	t.Skip("not implemented yet")
}

func TestRmUser(t *testing.T) {
	t.Skip("not implemented yet")
}
