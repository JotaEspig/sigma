package classroom

import (
	"errors"
	"sigma/models/student"

	"gorm.io/gorm"
)

type Classroom struct {
	gorm.Model
	Name     string
	Year     uint8
	Students []*student.Student
}

func InitClassroom(name string, year uint8) (*Classroom, error) {
	if name == "" {
		return nil, errors.New("classroom: Name cannot be empty")
	}

	c := &Classroom{
		Name: name,
		Year: year,
	}

	return c, nil
}

// Function that returns a map containing classroom info
func (c *Classroom) ToMap() map[string]interface{} {
	classroomMap := make(map[string]interface{})
	classroomMap["id"] = c.ID
	classroomMap["name"] = c.Name
	classroomMap["year"] = c.Year
	classroomMap["students"] = []map[string]interface{}{}
	for _, s := range c.Students {
		classroomMap["students"] = append(classroomMap["students"].([]map[string]interface{}), s.ToMap())
	}

	return classroomMap
}
