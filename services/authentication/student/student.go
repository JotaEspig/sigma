package studentauth

import (
	"database/sql"
	"errors"
	userauth "sigma/services/authentication/user"
)

type Student struct {
	ID      int
	User    *userauth.User
	Year    sql.NullInt16
	Status  sql.NullString
	ClassID sql.NullInt64 `db:"class_id"`
}

func InitStudent(u *userauth.User) (*Student, error) {
	if u.ID == 0 {
		return nil, errors.New("student: ID cannot be 0")
	}

	s := &Student{
		ID: u.ID,
	}

	return s, nil
}

// Adds student value to map contaning user info
func (s *Student) ToMap() map[string]interface{} {
	studentMap := s.User.ToMap()
	studentMap["year"] = s.Year.Int16
	studentMap["status"] = s.Status.String
	studentMap["class_id"] = s.ClassID.Int64
	return studentMap
}
