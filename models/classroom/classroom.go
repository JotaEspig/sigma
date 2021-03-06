package classroom

import (
	"errors"
	"sigma/models/student"
)

type Classroom struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"not null"`
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

	return classroomMap
}
