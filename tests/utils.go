package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// Constants to use in tests, def = default
const (
	defUsername      = "TestUsername"
	defPasswd        = "TestPasswd"
	defName          = "TestName"
	defSurname       = "TestSurname"
	defEmail         = "TestEmail"
	defClassroomName = "TestClassroom"
	defClassroomYear = 1111
)

// logs in and gets the token
func getToken(router *gin.Engine, username, password string) (string, bool) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/login",
		bytes.NewBuffer([]byte("username="+username+"&password="+password)),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	jsonResp := map[string]string{}
	json.Unmarshal(w.Body.Bytes(), &jsonResp)

	token, ok := jsonResp["token"]
	return token, ok

}
