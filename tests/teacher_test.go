package tests

import (
	"sigma/config"
	"sigma/models/teacher"
	"sigma/models/user"
	"testing"
)

func TestAddTeacher(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds user to be able to use teacher
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	teach, err := teacher.InitTeacher(u)
	if err != nil {
		t.Errorf("initializing teacher: %s", err)
	}

	// Adds teacher without education
	err = teacher.AddTeacher(config.DB, teach)
	if err != nil {
		t.Errorf("adding legit teacher (without education): %s", err)
	}

	// repeating the same action to check 'unique' columns
	err = teacher.AddTeacher(config.DB, teach)
	if err == nil {
		t.Errorf("adding repeated teacher (it's not supposed to happen): %s", err)
	}

	err = teacher.RmTeacher(config.DB, teach.User.Username)
	if err != nil {
		t.Errorf("removing legit teacher: %s", err)
	}

	teach.Education = "Bacharelado"

	// Adds teacher with education
	err = teacher.AddTeacher(config.DB, teach)
	if err != nil {
		t.Errorf("adding legit teacher: %s", err)
	}

	err = teacher.RmTeacher(config.DB, teach.User.Username)
	if err != nil {
		t.Errorf("removing legit teacher: %s", err)
	}

	err = user.RmUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestGetTeacher(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds user to be able to use teacher
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	teach, err := teacher.InitTeacher(u)
	if err != nil {
		t.Errorf("initializing teacher: %s", err)
	}

	teach.Education = "Bacharelado"

	// Adds teacher with education
	err = teacher.AddTeacher(config.DB, teach)
	if err != nil {
		t.Errorf("adding legit teacher (with education): %s", err)
	}

	// Gets teacher from database
	teach, err = teacher.GetTeacher(config.DB, defUsername)
	if err != nil {
		t.Errorf("getting legit teacher: %s", err)
	}

	// Checks if teacher is legit
	if teach.UID != u.ID {
		t.Errorf("teacher is not legit")
	}

	// Gets teacher from database with parcial info
	teach, err = teacher.GetTeacher(config.DB, defUsername, "education")
	if err != nil {
		t.Errorf("getting legit teacher: %s", err)
	}

	// Checks if teacher parcial info is legit
	if teach.Education != "Bacharelado" {
		t.Errorf("teacher parcial info is not legit")
	}

	// Gets non-existent teacher
	_, err = teacher.GetTeacher(config.DB, "non-existent-user")
	if err == nil {
		t.Errorf("getting non-existent teacher: %s", err)
	}

	err = teacher.RmTeacher(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit teacher: %s", err)
	}

	err = user.RmUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestUpdateTeacher(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds user to be able to use teacher
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	teach, err := teacher.InitTeacher(u)
	if err != nil {
		t.Errorf("initializing teacher: %s", err)
	}

	teach.Education = "Bacharelado"

	// Adds teacher with education
	err = teacher.AddTeacher(config.DB, teach)
	if err != nil {
		t.Errorf("adding legit teacher (with education): %s", err)
	}

	// Updates teacher
	teach.Education = "Mestrado"
	err = teacher.UpdateTeacher(config.DB, teach)
	if err != nil {
		t.Errorf("updating legit teacher: %s", err)
	}

	// Gets teacher from database
	teach, err = teacher.GetTeacher(config.DB, defUsername)
	if err != nil {
		t.Errorf("getting legit teacher: %s", err)
	}

	// Checks if teacher is legit
	if teach.UID != u.ID {
		t.Errorf("teacher is not legit")
	}

	// Checks if teacher is legit
	if teach.Education != "Mestrado" {
		t.Errorf("teacher is not legit")
	}

	// Gets teacher from database with parcial info
	teach, err = teacher.GetTeacher(config.DB, defUsername, "education")
	if err != nil {
		t.Errorf("getting legit teacher: %s", err)
	}

	// Checks if teacher parcial info is legit
	if teach.Education != "Mestrado" {
		t.Errorf("teacher parcial info is not legit")
	}

	err = teacher.RmTeacher(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit teacher: %s", err)
	}

	err = user.RmUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}
