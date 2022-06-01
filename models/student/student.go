package student

import (
	"errors"
	"sigma/models/classroom"
	"sigma/models/user"
)

type Student struct {
	UID         uint `gorm:"primary_key;column:id"`
	Status      string
	ClassroomID uint       `gorm:"default:null"`
	User        *user.User `gorm:"foreignKey:UID"`
	Classroom   *classroom.Classroom
}

func InitStudent(u *user.User) (*Student, error) {
	if u.ID == 0 {
		return nil, errors.New("student: UserID cannot be 0")
	}

	s := &Student{
		UID:  u.ID,
		User: u,
	}

	return s, nil
}

// Adds student value to map contaning user info
func (s *Student) ToMap() map[string]interface{} {
	studentMap := make(map[string]interface{})
	if s.User != nil {
		studentMap = s.User.ToMap()
	}

	studentMap["status"] = s.Status
	studentMap["classroom_id"] = s.ClassroomID
	if s.Classroom != nil {
		studentMap["classroom"] = s.Classroom.ToMap()
	}

	return studentMap
}
