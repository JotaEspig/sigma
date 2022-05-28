package tests

import (
	"sigma/db"
	"sigma/models/student"
	"sigma/models/user"
	"testing"
)

func TestAddStudent(t *testing.T) {
	db.DB.AutoMigrate(&student.Student{})

	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds user to be able to use student
	user.AddUser(db.DB, u)

	u, err := user.GetUser(db.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	s, err := student.InitStudent(u)
	if err != nil {
		t.Errorf("initializing student: %s", err)
	}

	// Adds student without year and status
	err = student.AddStudent(db.DB, s)
	if err != nil {
		t.Errorf("adding legit user (without year and status): %s", err)
	}

	// repeating the same action to check 'unique' columns
	err = student.AddStudent(db.DB, s)
	if err == nil {
		t.Errorf("adding repeated student (it's not supposed to happen): %s", err)
	}

	err = student.RmStudent(db.DB, s.User.Username)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}

	s.Status = "ativo"

	// Adds user with year and status
	err = student.AddStudent(db.DB, s)
	if err != nil {
		t.Errorf("adding legit user (without year and status): %s", err)
	}

	err = student.RmStudent(db.DB, s.User.Username)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}

	err = student.RmStudent(db.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit student: %s", err)
	}
	err = user.RmUser(db.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestGetStudent(t *testing.T) {
	db.DB.AutoMigrate(&student.Student{})
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	user.AddUser(db.DB, u)

	u, err := user.GetUser(db.DB, defUsername)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	s, err := student.InitStudent(u)
	if err != nil {
		t.Errorf("initializing student: %s", err)
	}

	s.Status = "ativo"

	// Adds user with year and status
	err = student.AddStudent(db.DB, s)
	if err != nil {
		t.Errorf("adding legit user (without year and status): %s", err)
	}

	_, err = student.GetStudent(db.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit student: %s", err)
	}

	// Checks if get student parcial info is working
	s, err = student.GetStudent(db.DB, u.Username, "year", "status")
	if err != nil {
		t.Errorf("getting legit student (parcial info): %s", err)
	}
	if s.Status == "" {
		t.Errorf("getting legit student (parcial info): status is empty")
	}

	_, err = student.GetStudent(db.DB, "non-existent-username")
	if err == nil {
		t.Errorf("getting non existent student (it's not supposed to work): %s", err)
	}

	err = student.RmStudent(db.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit student: %s", err)
	}
	err = user.RmUser(db.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestRmStudent(t *testing.T) {
	db.DB.AutoMigrate(&user.User{})

	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	user.AddUser(db.DB, u)

	err := student.RmStudent(db.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit student: %s", err)
	}
	err = user.RmUser(db.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}
