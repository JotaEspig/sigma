package student

import (
	dbPKG "sigma/db"
	"sigma/models/user"

	"gorm.io/gorm"
)

// Slice of all student params
var StudentParams = []string{
	"id",
	"year",
	"status",
	"user_id",
}

// Slice of public student params
var PublicStudentParams = []string{
	"id",
	"year",
	"user_id",
}

// Adds a student to a database.
func AddStudent(db *gorm.DB, s *Student) error {
	return db.Transaction(func(tx *gorm.DB) error {
		s.User.Type = "student"
		err := tx.Save(s.User).Error
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
func GetStudent(db *gorm.DB, username string, columns ...string) (*Student, error) {
	s := &Student{}

	u, err := user.GetUser(db, username, "id")
	if err != nil {
		return nil, err
	}

	columnsToUse := dbPKG.GetColumns(StudentParams, columns...)

	err = db.Select(columnsToUse).Where("user_id = ?", u.ID).First(s).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(s).Association("User").Find(&s.User)

	return s, err
}

// Removes a student from a database
func RmStudent(db *gorm.DB, username string) error {
	u, err := user.GetUser(db, username, "id")
	if err != nil {
		return err
	}

	return db.Unscoped().Where("user_id = ?", u.ID).Delete(&Student{}).Error
}

// AutoMigrate the student table
func init() {
	dbPKG.DB.AutoMigrate(&Student{})
}
