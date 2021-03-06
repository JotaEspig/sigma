package teacher

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
	dbPKG "sigma/db"
	"sigma/models/user"

	"gorm.io/gorm"
)

// Slice of all teacher params
var TeacherParams = []string{
	"id",
	"education",
}

// Slice of public teacher params
var PublicTeacherParams = []string{
	"id",
	"education",
}

// Adds a teacher to a database.
func AddTeacher(db *gorm.DB, t *Teacher) error {
	return db.Transaction(func(tx *gorm.DB) error {
		t.User.Type = "teacher"
		err := user.UpdateUser(db, t.User)
		if err != nil {
			return err
		}

		err = tx.Create(t).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// Gets a teacher from a database
func GetTeacher(db *gorm.DB, username string, params ...string) (*Teacher, error) {
	t := &Teacher{}

	u, err := user.GetUser(db, username, "id")
	if err != nil {
		return nil, err
	}

	columnsToUse := dbPKG.GetColumns(TeacherParams, params...)

	err = db.Select(columnsToUse).Where("id = ?", u.ID).First(t).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(t).Association("User").Find(&t.User)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Updates a teacher in the database
func UpdateTeacher(db *gorm.DB, t *Teacher) error {
	return db.Model(t).Updates(t).Error
}

// Deletes a teacher from the database
func RmTeacher(db *gorm.DB, username string) error {
	u, err := user.GetUser(db, username, "id")
	if err != nil {
		return err
	}

	return db.Delete(&Teacher{UID: u.ID}).Error
}

func init() {
	config.DB.AutoMigrate(&Teacher{})
}
