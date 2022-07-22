package tests

import (
	"sigma/config"
	"sigma/models/student"
	"sigma/models/user"
	"testing"
)

func TestAddStudent(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds user to be able to use student
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	s, err := student.InitStudent(u)
	if err != nil {
		t.Errorf("initializing student: %s", err)
	}

	// Adds student without status
	err = student.AddStudent(config.DB, s)
	if err != nil {
		t.Errorf("adding legit user (without status): %s", err)
	}

	// repeating the same action to check 'unique' columns
	err = student.AddStudent(config.DB, s)
	if err == nil {
		t.Errorf("adding repeated student (it's not supposed to happen): %s", err)
	}

	err = student.RmStudent(config.DB, s.User.Username)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}

	s.Status = "ativo"

	// Adds user with status
	err = student.AddStudent(config.DB, s)
	if err != nil {
		t.Errorf("adding legit user: %s", err)
	}

	err = student.RmStudent(config.DB, s.User.Username)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}

	err = student.RmStudent(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit student: %s", err)
	}
	err = user.RmUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestGetStudent(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	s, err := student.InitStudent(u)
	if err != nil {
		t.Errorf("initializing student: %s", err)
	}

	s.Status = "ativo"

	// Adds user with year and status
	err = student.AddStudent(config.DB, s)
	if err != nil {
		t.Errorf("adding legit user (with status): %s", err)
	}

	_, err = student.GetStudent(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit student: %s", err)
	}

	// Checks if get student parcial info is working
	s, err = student.GetStudent(config.DB, u.Username, "status")
	if err != nil {
		t.Errorf("getting legit student (parcial info): %s", err)
	}
	if s.Status == "" {
		t.Errorf("getting legit student (parcial info): status is empty")
	}

	// Gets non-existent student
	_, err = student.GetStudent(config.DB, "non-existent-username")
	if err == nil {
		t.Errorf("getting non existent student (it's not supposed to work): %s", err)
	}

	err = student.RmStudent(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit student: %s", err)
	}
	err = user.RmUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestUpdateStudent(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	s, err := student.InitStudent(u)
	if err != nil {
		t.Errorf("initializing student: %s", err)
	}

	s.Status = "ativo"

	// Adds user with year and status
	err = student.AddStudent(config.DB, s)
	if err != nil {
		t.Errorf("adding legit user (with status): %s", err)
	}

	s.Status = "inativo"

	err = student.UpdateStudent(config.DB, s)
	if err != nil {
		t.Errorf("updating legit student: %s", err)
	}

	s, err = student.GetStudent(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit student: %s", err)
	}
	if s.Status != "inativo" {
		t.Errorf("updating legit student: status is not inativo")
	}

	err = student.RmStudent(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit student: %s", err)
	}
	err = user.RmUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestRmStudent(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds if user's not added in the database
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	s, err := student.InitStudent(u)
	if err != nil {
		t.Errorf("initializing student: %s", err)
	}

	err = student.AddStudent(config.DB, s)
	if err != nil {
		t.Errorf("adding legit student: %s", err)
	}

	err = student.RmStudent(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit student: %s", err)
	}
	err = user.RmUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}
