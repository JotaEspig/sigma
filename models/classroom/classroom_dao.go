package classroom

import (
	// dbPKG means 'the package db', because if it's named db
	// it will conflict with db variable in the functions below
	"sigma/config"
	dbPKG "sigma/db"

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

// Adds a classroom to a database
func AddClassroom(db *gorm.DB, c *Classroom) error {
	return db.Create(c).Error
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
		err = db.Model(&c.Students[i]).Association("User").Find(&c.Students[i].User)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Gets all classroom from a database
func GetAllClassrooms(db *gorm.DB, params ...string) ([]Classroom, error) {
	classrooms := []Classroom{}

	columnsToUse := dbPKG.GetColumns(ClassroomParams, params...)

	err := db.Select(columnsToUse).Find(&classrooms).Error
	if err != nil {
		return classrooms, err
	}

	return classrooms, nil
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
