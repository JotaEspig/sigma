package tests

import (
	"sigma/config"
	"sigma/models/classroom"
	"sigma/models/student"
	"sigma/models/user"
	"testing"
)

const (
	defClassroomName = "TestClassroom"
	defClassroomYear = 1
)

func TestAddClassroom(t *testing.T) {
	c, err := classroom.InitClassroom(defClassroomName, defClassroomYear)
	if err != nil {
		t.Errorf("initializing classroom: %s", err)
	}

	err = classroom.AddClassroom(config.DB, c)
	if err != nil {
		t.Errorf("adding legit classroom: %s", err)
	}

	// repeating the same action to check 'unique' columns
	err = classroom.AddClassroom(config.DB, c)
	if err == nil {
		t.Errorf("adding repeated classroom (it's not supposed to happen): %s", err)
	}

	err = classroom.RmClassroom(config.DB, c.ID)
	if err != nil {
		t.Errorf("removing legit classroom: %s", err)
	}
}

func TestGetClassroom(t *testing.T) {
	c, err := classroom.InitClassroom(defClassroomName, defClassroomYear)
	if err != nil {
		t.Errorf("initializing classroom: %s", err)
	}

	err = classroom.AddClassroom(config.DB, c)
	if err != nil {
		t.Errorf("adding legit classroom: %s", err)
	}

	c, err = classroom.GetClassroom(config.DB, c.ID)
	if err != nil {
		t.Errorf("getting legit classroom: %s", err)
	}
	if c.ID == 0 {
		t.Errorf("getting legit classroom: classroom id is null")
	}

	// adding students to classroom
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)
	err = user.AddUser(config.DB, u)
	if err != nil {
		t.Errorf("adding legit user: %s", err)
	}

	s, err := student.InitStudent(u)
	if err != nil {
		t.Errorf("initializing student: %s", err)
	}

	s.ClassroomID = c.ID
	err = student.AddStudent(config.DB, s)
	if err != nil {
		t.Errorf("adding legit student: %s", err)
	}

	c, err = classroom.GetClassroom(config.DB, c.ID)
	if err != nil {
		t.Errorf("getting legit classroom: %s", err)
	}
	if len(c.Students) != 1 {
		t.Errorf("getting legit classroom: students count is not 1")
	}

	// removing students and users after test
	err = student.RmStudent(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit student: %s", err)
	}

	err = user.RmUser(config.DB, defUsername)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}

	// Gets parcial info of classroom
	c, err = classroom.GetClassroom(config.DB, c.ID, "id", "name")
	if err != nil {
		t.Errorf("getting parcial info of classroom: %s", err)
	}
	if c.Name == "" {
		t.Errorf("getting parcial info of classroom: classroom name is empty")
	}

	// Gets non-existent classroom
	_, err = classroom.GetClassroom(config.DB, 0)
	if err == nil {
		t.Errorf("getting non-existent classroom (it's not supposed to work): %s", err)
	}

	err = classroom.RmClassroom(config.DB, c.ID)
	if err != nil {
		t.Errorf("removing legit classroom: %s", err)
	}
}

func TestGetAllClassrooms(t *testing.T) {
	c1, err := classroom.InitClassroom(defClassroomName, defClassroomYear)
	if err != nil {
		t.Errorf("initializing classroom: %s", err)
	}

	err = classroom.AddClassroom(config.DB, c1)
	if err != nil {
		t.Errorf("adding legit classroom: %s", err)
	}

	c2, err := classroom.InitClassroom("other classroom name", defClassroomYear)
	if err != nil {
		t.Errorf("initializing classroom: %s", err)
	}

	err = classroom.AddClassroom(config.DB, c2)
	if err != nil {
		t.Errorf("adding legit classroom: %s", err)
	}

	cs, err := classroom.GetAllClassrooms(config.DB)
	if err != nil {
		t.Errorf("getting all classrooms: %s", err)
	}
	if len(cs) != 2 {
		t.Errorf("getting all classrooms: classrooms count is 0")
	}

	err = classroom.RmClassroom(config.DB, c1.ID)
	if err != nil {
		t.Errorf("removing legit classroom: %s", err)
	}

	err = classroom.RmClassroom(config.DB, c2.ID)
	if err != nil {
		t.Errorf("removing legit classroom: %s", err)
	}
}

func TestUpdateClassroom(t *testing.T) {
	c, err := classroom.InitClassroom(defClassroomName, defClassroomYear)
	if err != nil {
		t.Errorf("initializing classroom: %s", err)
	}

	err = classroom.AddClassroom(config.DB, c)
	if err != nil {
		t.Errorf("adding legit classroom: %s", err)
	}

	c.Name = "different classroom name"
	err = classroom.UpdateClassroom(config.DB, c)
	if err != nil {
		t.Errorf("updating legit classroom: %s", err)
	}

	c, err = classroom.GetClassroom(config.DB, c.ID)
	if err != nil {
		t.Errorf("getting legit classroom: %s", err)
	}
	if c.Name != "different classroom name" {
		t.Errorf("updating legit classroom: classroom name is not updated")
	}

	err = classroom.RmClassroom(config.DB, c.ID)
	if err != nil {
		t.Errorf("removing legit classroom: %s", err)
	}

}

func TestRmClassroom(t *testing.T) {
	c, err := classroom.InitClassroom(defClassroomName, defClassroomYear)
	if err != nil {
		t.Errorf("initializing classroom: %s", err)
	}

	err = classroom.AddClassroom(config.DB, c)
	if err != nil {
		t.Errorf("adding legit classroom: %s", err)
	}

	err = classroom.RmClassroom(config.DB, c.ID)
	if err != nil {
		t.Errorf("removing legit classroom: %s", err)
	}
}
