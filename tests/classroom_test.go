package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sigma/config"
	"sigma/models/classroom"
	"sigma/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddClassroom(t *testing.T) {
	router := server.CreateTestRouter()

	token, ok := getToken(router, "admin", "admin")
	assert.Equal(t, true, ok)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/admin/tools/classroom/add",
		bytes.NewBuffer([]byte(
			`{"name": `+defClassroomName+`, "year": `+fmt.Sprint(defClassroomYear)+`}`,
		)),
	)
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: token,
	})

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	err := config.DB.Unscoped().Delete(&classroom.Classroom{}, "name = ?", defClassroomName).Error
	assert.Equal(t, nil, err)
}

func TestGetClassroom(t *testing.T) {
	router := server.CreateTestRouter()

	token, ok := getToken(router, "admin", "admin")
	assert.Equal(t, true, ok)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/admin/tools/classroom/add",
		bytes.NewBuffer([]byte(
			`{"name": `+defClassroomName+`, "year": `+fmt.Sprint(defClassroomYear)+`}`,
		)),
	)
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: token,
	})

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/admin/tools/classroom/get", nil)
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: token,
	})

	err := config.DB.Unscoped().Delete(&classroom.Classroom{}, "name = ?", "test").Error
	assert.Equal(t, nil, err)

	// Test GetClassroom endpoint with http.NewRequest
	t.Skip("Skipping GetClassroom endpoint test, because it's not implemented yet")
}

func TestGetAllClassrooms(t *testing.T) {
	_ = server.CreateRouter()

	// Test GetAllClassrooms endpoint with http.NewRequest
	t.Skip("Skipping GetAllClassrooms endpoint test, because it's not implemented yet")
}
