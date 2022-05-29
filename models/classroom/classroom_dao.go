package classroom

import (
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

	err := db.Select(columnsToUse).Where("id = ?", id).First(c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Updates a classroom in a database
func UpdateClassroom(db *gorm.DB, c *Classroom) error {
	return db.Save(c).Error
}

// Removes a classroom from a database
func RmClassroom(db *gorm.DB, id uint) error {
	return db.Unscoped().Delete(&Classroom{}, "id = ?", id).Error
}

func init() {
	dbPKG.DB.AutoMigrate(&Classroom{})
}