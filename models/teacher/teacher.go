package teacher

import (
	"errors"
	"sigma/models/user"
)

// Teacher represents a teacher role in sigma
type Teacher struct {
	UID       uint `gorm:"primary_key;column:id"`
	Education string
	User      *user.User `gorm:"foreignKey:UID"`
}

// InitTeacher initializes a teacher struct
func InitTeacher(u *user.User) (*Teacher, error) {
	if u.ID == 0 {
		return nil, errors.New("teacher: UserID cannot be 0")
	}

	t := &Teacher{
		UID:  u.ID,
		User: u,
	}

	return t, nil
}

// ToMap adds teacher value to map contaning user info
func (t Teacher) ToMap() map[string]interface{} {
	teacherMap := make(map[string]interface{})
	if t.User != nil {
		teacherMap = t.User.ToMap()
	}
	teacherMap["education"] = t.Education

	return teacherMap
}
