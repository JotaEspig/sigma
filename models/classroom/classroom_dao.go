package classroom

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
	dbPKG "sigma/db"
	"sigma/models/user"

	"gorm.io/gorm"
)

// Slice containing all classroom params
var ClassroomParams = []string{
	"id",
	"name",
	"year",
}

// Slice containing public classroom params
var PublicClassroomParams = []string{
	"id",
	"name",
	"year",
}

// Gets a classroom from a database
func GetClassroom(db *gorm.DB, id uint, params ...string) (*Classroom, error) {
	c := &Classroom{}

	columnsToUse := dbPKG.GetColumns(ClassroomParams, params...)

	err := db.Select(columnsToUse).Where("id = ?", id).Preload("Students").First(c).Error
	if err != nil {
		return nil, err
	}

	for i := range c.Students {
		user_id := c.Students[i].UID
		// Loads parcial user data for each student to avoid loading the whole user
		err = db.Model(&user.User{}).Select("id", "name", "surname").Where("id = ?", user_id).
			First(&c.Students[i].User).Error

		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Updates a classroom in a database
func UpdateClassroom(db *gorm.DB, c *Classroom) error {
	return db.Model(c).Updates(c).Error
}

// Removes a classroom from a database
func RmClassroom(db *gorm.DB, id uint) error {
	return db.Unscoped().Delete(&Classroom{}, "id = ?", id).Error
}

func init() {
	config.DB.AutoMigrate(&Classroom{})
}
