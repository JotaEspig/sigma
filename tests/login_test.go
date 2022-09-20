package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os/user"
	"sigma/config"
	"sigma/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	router := server.CreateTestRouter()
	_, ok := getToken(router, "admin", "admin")
	assert.Equal(t, true, ok)
}

func TestSignup(t *testing.T) {
	router := server.CreateTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/cadastro",
		bytes.NewBuffer([]byte(
			fmt.Sprintf(
				"name=%s&surname=%s&username=%s&email=%s&password=%s",
				defName, defSurname, defUsername, defEmail, defPasswd,
			),
		)),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	config.DB.Unscoped().Where("username = ?", defUsername).Delete(&user.User{})
}
