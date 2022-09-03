package teacher

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
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

// Default function to update a teacher in the database
func UpdateTeacher(db *gorm.DB, t *Teacher) error {
	return db.Model(t).Omit("id").Updates(t).Error
}

// Deletes a teacher from the database
func RmTeacher(db *gorm.DB, username string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		u := user.User{}
		err := config.DB.Select("id").Where("username = ?", username).First(&u).Error
		if err != nil {
			return err
		}

		// Updates the type of the user to be empty
		err = db.Model(u).Update("type", "").Error
		if err != nil {
			return err
		}

		return db.Unscoped().Delete(&Teacher{}, "id = ?", u.ID).Error
	})
}

func init() {
	config.DB.AutoMigrate(&Teacher{})
}
