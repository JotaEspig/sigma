package student

import (
	"errors"
	"sigma/models/user"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Status string
	UserID uint `gorm:"not null;unique"`
	User   *user.User
}

func InitStudent(u *user.User) (*Student, error) {
	if u.ID == 0 {
		return nil, errors.New("student: UserID cannot be 0")
	}

	s := &Student{
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
	studentMap["user"] = s.User.ToMap()
	return studentMap
}
