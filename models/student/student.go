package student

import (
	"errors"
	"sigma/models/classroom"
	"sigma/models/user"
)

type Student struct {
	ID          uint `gorm:"primary_key"`
	Status      string
	UserID      uint `gorm:"not null;unique"`
	ClassroomID uint `gorm:"default:null"`
	User        *user.User
	Classroom   *classroom.Classroom
}

func InitStudent(u *user.User) (*Student, error) {
	if u.ID == 0 {
		return nil, errors.New("student: UserID cannot be 0")
	}

	s := &Student{
		ID:     u.ID,
		UserID: u.ID,
		User:   u,
	}

	return s, nil
}

// Adds student value to map contaning user info
func (s *Student) ToMap() map[string]interface{} {
	studentMap := s.User.ToMap()
	studentMap["status"] = s.Status
	studentMap["user_id"] = s.UserID
	studentMap["classroom_id"] = s.ClassroomID
	studentMap["classroom"] = s.Classroom.ToMap()
	return studentMap
}
