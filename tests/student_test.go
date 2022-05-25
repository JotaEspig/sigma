package tests

import (
	"sigma/db"
	"sigma/models/user"
	"testing"
)

func TestAddStudent(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	user.AddUser(db.DB, u)
	u, err := user.GetUser(db.DB, u.Username)
	if err != nil {
		t.Error(err)
	}

	s, err := InitStudent(u)
	if err != nil {
		t.Error(err)
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("adding legit student: %s", r)
			}
		}()
		user.AddStudent(db.DB, s)
	}()

	// repeating the same action
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("adding repeated student (it's not supposed to work)")
			}
		}()
		user.AddStudent(db.DB, s)
	}()
}

func TestGetStudent(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		user.AddUser(db.DB, u)
	}()

	u, err := user.GetUser(db.DB, defUsername)
	if err != nil {
		t.Errorf("getting legit student: %s", err)
	}

	s, err := InitStudent(u)
	if err != nil {
		t.Error(err)
	}

	s.Year.Scan(2)
	s.Status.Scan("ativo")

	// Adds student
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Error(r)
			}
		}()
		user.AddStudent(db.DB, s)
	}()

	_, err = GetStudent(db.DB, u.ID)
	// Checks if get student parcial info is working
	if err != nil {
		t.Errorf("getting legit student (parcial info): %s", err)
	}

	s, err = GetStudent(db.DB, u.ID, "year", "status")
	if err != nil {
		t.Errorf("getting legit student (parcial info): %s", err)
	}
	if s.Year.Int16 == 0 {
		t.Errorf("getting legit student (parcial info): year is empty")
	}
	if s.Status.String == "" {
		t.Errorf("getting legit student (parcial info): status is empty")
	}
	if s.ClassID.Int64 != 0 {
		t.Errorf("getting legit student (parcial info): class id is filled")
	}

	_, err = GetStudent(db.DB, nonExistentID)
	if err == nil {
		t.Errorf("getting non existent student (it's not supposed to work): %s", err)
	}
}
