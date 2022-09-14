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
	"strings"
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
			"name="+strings.ReplaceAll(defClassroomName, " ", "+")+
				"&year="+fmt.Sprint(defClassroomYear),
		)),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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

	c, _ := classroom.InitClassroom(defClassroomName, defClassroomYear)
	err := config.DB.Create(c).Error
	assert.Equal(t, nil, err)

	token, ok := getToken(router, "admin", "admin")
	assert.Equal(t, true, ok)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/tools/classroom/get", nil)
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: token,
	})

	router.ServeHTTP(w, req)
	classrooms := []classroom.Classroom{}
	json.Unmarshal(w.Body.Bytes(), &classrooms)

	classID := 0
	for _, c := range classrooms {
		if c.Name == defClassroomName && c.Year == defClassroomYear {
			classID = int(c.ID)
			break
		}
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(
		"GET",
		"/admin/tools/classroom/"+fmt.Sprint(classID)+"/get",
		nil,
	)
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: token,
	})

	router.ServeHTTP(w, req)
	cNew := classroom.Classroom{}
	json.Unmarshal(w.Body.Bytes(), &cNew)
	assert.Equal(t, defClassroomName, cNew.Name)
	assert.Equal(t, defClassroomYear, cNew.Year)

	err = config.DB.Unscoped().Delete(&classroom.Classroom{}, "name = ?", defClassroomName).Error
	assert.Equal(t, nil, err)
}

func TestGetAllClassrooms(t *testing.T) {
	router := server.CreateTestRouter()

	c, _ := classroom.InitClassroom(defClassroomName, defClassroomYear)
	err := config.DB.Create(c).Error
	assert.Equal(t, nil, err)

	token, ok := getToken(router, "admin", "admin")
	assert.Equal(t, true, ok)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/tools/classroom/get", nil)
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: token,
	})

	router.ServeHTTP(w, req)
	classrooms := []classroom.Classroom{}
	json.Unmarshal(w.Body.Bytes(), &classrooms)
	assert.GreaterOrEqual(t, len(classrooms), 1)

	err = config.DB.Unscoped().Delete(&classroom.Classroom{}, "name = ?", defClassroomName).Error
	assert.Equal(t, nil, err)
}
