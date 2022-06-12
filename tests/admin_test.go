package tests

import (
	"sigma/config"
	"sigma/models/admin"
	"sigma/models/user"
	"testing"
)

func TestAddAdmin(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds user to be able to use admin
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	a, err := admin.InitAdmin(u)
	if err != nil {
		t.Errorf("initializing admin: %s", err)
	}

	// Adds admin without role
	err = admin.AddAdmin(config.DB, a)
	if err != nil {
		t.Errorf("adding legit admin (without role): %s", err)
	}

	// repeating the same action to check 'unique' columns
	err = admin.AddAdmin(config.DB, a)
	if err == nil {
		t.Errorf("adding repeated admin (it's not supposed to happen): %s", err)
	}

	err = admin.RmAdmin(config.DB, a.User.Username)
	if err != nil {
		t.Errorf("removing legit admin: %s", err)
	}

	a.Role = "admin"

	// Adds admin with role
	err = admin.AddAdmin(config.DB, a)
	if err != nil {
		t.Errorf("adding legit admin: %s", err)
	}

	err = admin.RmAdmin(config.DB, a.User.Username)
	if err != nil {
		t.Errorf("removing legit admin: %s", err)
	}

	err = user.RmUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestGetAdmin(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds user to be able to use admin
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	a, err := admin.InitAdmin(u)
	if err != nil {
		t.Errorf("initializing admin: %s", err)
	}

	// Adds admin without role
	err = admin.AddAdmin(config.DB, a)
	if err != nil {
		t.Errorf("adding legit admin (without role): %s", err)
	}

	err = admin.RmAdmin(config.DB, a.User.Username)
	if err != nil {
		t.Errorf("removing legit admin: %s", err)
	}

	a.Role = "coordenador"

	// Adds admin with role
	err = admin.AddAdmin(config.DB, a)
	if err != nil {
		t.Errorf("adding legit admin: %s", err)
	}

	a, err = admin.GetAdmin(config.DB, a.User.Username)
	if err != nil {
		t.Errorf("getting legit admin: %s", err)
	}

	err = admin.RmAdmin(config.DB, a.User.Username)
	if err != nil {
		t.Errorf("removing legit admin: %s", err)
	}

	err = user.RmUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestUpdateAdmin(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds user to be able to use admin
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	a, err := admin.InitAdmin(u)
	if err != nil {
		t.Errorf("initializing admin: %s", err)
	}

	err = admin.AddAdmin(config.DB, a)
	if err != nil {
		t.Errorf("adding legit admin: %s", err)
	}

	a.Role = "coordenador"

	err = admin.UpdateAdmin(config.DB, a)
	if err != nil {
		t.Errorf("updating legit admin: %s", err)
	}

	err = admin.RmAdmin(config.DB, a.User.Username)
	if err != nil {
		t.Errorf("removing legit admin: %s", err)
	}

	err = user.RmUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}

func TestRmAdmin(t *testing.T) {
	u := user.InitUser(defUsername, defEmail, defName, defSurname, defPasswd)

	// Adds user to be able to use admin
	user.AddUser(config.DB, u)

	u, err := user.GetUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("getting legit user: %s", err)
	}

	a, err := admin.InitAdmin(u)
	if err != nil {
		t.Errorf("initializing admin: %s", err)
	}

	err = admin.AddAdmin(config.DB, a)
	if err != nil {
		t.Errorf("adding legit admin (without role): %s", err)
	}

	err = admin.RmAdmin(config.DB, a.User.Username)
	if err != nil {
		t.Errorf("removing legit admin: %s", err)
	}

	err = user.RmUser(config.DB, u.Username)
	if err != nil {
		t.Errorf("removing legit user: %s", err)
	}
}
