package tests

import (
	"sigma/config"
	"sigma/models/classroom"
	"testing"
)

const (
	defClassroomName = "TestClassroom"
	defClassroomYear = 1
)

func TestAddClassroom(t *testing.T) {
	config.DB.AutoMigrate(&classroom.Classroom{})

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
	config.DB.AutoMigrate(&classroom.Classroom{})

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

	// Gets parcial info of classroom
	c, err = classroom.GetClassroom(config.DB, c.ID, "name")
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

func TestUpdateClassroom(t *testing.T) {
	config.DB.AutoMigrate(&classroom.Classroom{})

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
	config.DB.AutoMigrate(&classroom.Classroom{})

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
