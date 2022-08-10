package student

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
	dbPKG "sigma/db"
	"sigma/models/user"

	"gorm.io/gorm"
)

// Slice of all student params
var StudentParams = []string{
	"id",
	"status",
}

// Slice of public student params
var PublicStudentParams = []string{
	"id",
}

// Adds a student to a database.
func AddStudent(db *gorm.DB, s *Student) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := db.Model(s.User).Update("type", "student").Error
		if err != nil {
			return err
		}

		err = tx.Create(s).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// Gets a student from a database
func GetStudent(db *gorm.DB, username string, params ...string) (*Student, error) {
	s := &Student{}

	u, err := user.GetUser(db, username, "id")
	if err != nil {
		return nil, err
	}

	columnsToUse := dbPKG.GetColumns(StudentParams, params...)

	err = db.Select(columnsToUse).Where("id = ?", u.ID).First(s).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(s).Association("User").Find(&s.User)
	if err != nil {
		return nil, err
	}

	return s, err
}

// Default function to update a student in a database
func UpdateStudent(db *gorm.DB, s *Student) error {
	return db.Model(s).Omit("id").Updates(s).Error
}

// Removes a student from a database
func RmStudent(db *gorm.DB, username string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		u, err := user.GetUser(db, username, "id")
		if err != nil {
			return err
		}

		// Updates the type of the user to be empty
		err = db.Model(u).Update("type", "").Error
		if err != nil {
			return err
		}

		return db.Unscoped().Delete(&Student{}, "id = ?", u.ID).Error
	})
}

// AutoMigrate the student table
func init() {
	config.DB.AutoMigrate(&Student{})
}
