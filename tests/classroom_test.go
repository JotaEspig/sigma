package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sigma/config"
	"sigma/models/classroom"
	"sigma/models/student"
	"sigma/models/user"
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

	err = config.DB.Where("name = ? AND year = ?", defClassroomName, defClassroomYear).
		Select("id").First(c).Error
	assert.Equal(t, nil, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"GET",
		"/admin/tools/classroom/"+fmt.Sprint(c.ID)+"/get",
		nil,
	)
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: token,
	})

	router.ServeHTTP(w, req)
	jsonResponse := map[string]classroom.Classroom{
		"classroom": {},
	}
	json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.Equal(t, defClassroomName, jsonResponse["classroom"].Name)
	assert.Equal(t, defClassroomYear, jsonResponse["classroom"].Year)

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
	jsonResponse := map[string][]classroom.Classroom{
		"classrooms": {},
	}
	json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.GreaterOrEqual(t, len(jsonResponse["classrooms"]), 1)

	err = config.DB.Unscoped().Delete(&classroom.Classroom{}, "name = ?", defClassroomName).Error
	assert.Equal(t, nil, err)
}

func TestUpdateClassroom(t *testing.T) {
	router := server.CreateTestRouter()

	c, _ := classroom.InitClassroom(defClassroomName, defClassroomYear)
	err := config.DB.Create(c).Error
	assert.Equal(t, nil, err)

	token, ok := getToken(router, "admin", "admin")
	assert.Equal(t, true, ok)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"PUT",
		fmt.Sprintf("/admin/tools/classroom/%d/update", c.ID),
		bytes.NewBuffer([]byte(
			"name=TestUpdateClassroom&year=6969",
		)),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: token,
	})

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	err = config.DB.First(&c, c.ID).Error
	assert.Equal(t, nil, err)
	assert.Equal(t, "TestUpdateClassroom", c.Name)
	assert.Equal(t, uint16(6969), c.Year)

	err = config.DB.Unscoped().Delete(&c).Error
	assert.Equal(t, nil, err)
}

func TestDeleteClassroom(t *testing.T) {
	router := server.CreateTestRouter()

	c, _ := classroom.InitClassroom(defClassroomName, defClassroomYear)
	err := config.DB.Create(c).Error
	assert.Equal(t, nil, err)

	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	err = config.DB.Create(u).Error
	assert.Equal(t, nil, err)

	s, err := student.InitStudent(u)
	assert.Equal(t, nil, err)
	s.ClassroomID = c.ID
	err = config.DB.Create(s).Error
	assert.Equal(t, nil, err)

	err = config.DB.Preload("Students.User").First(c, c.ID).Error
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(c.Students))
	assert.Equal(t, defUsername, c.Students[0].User.Username)

	token, ok := getToken(router, "admin", "admin")
	assert.Equal(t, true, ok)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("/admin/tools/classroom/%d/delete", c.ID),
		nil,
	)
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: token,
	})

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	err = config.DB.Select("id", "classroom_id").First(s, s.UID).Error
	assert.Equal(t, nil, err)
	assert.Equal(t, uint(0), s.ClassroomID)

	err = config.DB.Unscoped().Delete(s).Error
	assert.Equal(t, nil, err)
	err = config.DB.Unscoped().Delete(u).Error
	assert.Equal(t, nil, err)
}
