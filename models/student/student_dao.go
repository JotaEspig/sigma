package student

import (
	dbPKG "sigma/db"
	"sigma/models/user"

	"gorm.io/gorm"
)

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

	columnsToUse := dbPKG.GetColumns(columns...)

	err = db.Select(columnsToUse).Where("user_id = ?", u.Model.ID).First(s).Error
	return s, err
}

// Removes a student from a database
func RmStudent(db *gorm.DB, username string) error {
	u, err := user.GetUser(db, username, "user_id")
	if err != nil {
		return err
	}

	return db.Unscoped().Where("username = ?", u.Username).Delete(&Student{}).Error
}
