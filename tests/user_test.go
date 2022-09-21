package tests

import (
	"net/http"
	"net/http/httptest"
	"sigma/config"
	"sigma/models/user"
	"sigma/server"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestSearchUsers(t *testing.T) {
	router := server.CreateTestRouter()

	usernForTest := "admin"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"GET",
		"/search/users/"+usernForTest,
		nil,
	)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetPublicUser(t *testing.T) {
	router := server.CreateTestRouter()

	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	config.DB.Create(u)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"GET",
		"/"+defUsername+"/get",
		nil,
	)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	config.DB.Delete(u)
}

func TestUpdateUser(t *testing.T) {
	t.Skip("not implemented yet")
}

func TestRmUser(t *testing.T) {
	t.Skip("not implemented yet")
}
