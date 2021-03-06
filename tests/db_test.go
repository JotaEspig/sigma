package tests

import (
	"sigma/db"
	"sigma/models/user"
	"testing"
)

func TestGetColumns(t *testing.T) {
	columns := []string{"username", "password"}

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Error(r)
			}
		}()

		newColumns := db.GetColumns(user.UserParams, columns...)
		if newColumns[0] != "username" {
			t.Errorf("get columns: There is no username in first index")
		}
		if newColumns[1] != "password" {
			t.Errorf("get columns: There is no password in seconde index")
		}
	}()

}
