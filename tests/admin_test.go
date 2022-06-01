package tests

import (
	"sigma/models/admin"
	"testing"
)

func TestAdmin(t *testing.T) {
	a := &admin.Admin{
		UID: 1,
	}
	if a.UID != 1 {
		t.Error("Admin.ID is not 1")
	}
}
