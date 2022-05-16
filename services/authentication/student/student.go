package studentauth

import (
	"database/sql"
	userauth "sigma/services/authentication/user"
)

type Student struct {
	User    *userauth.User
	Year    sql.NullInt16
	Status  sql.NullString
	ClassID sql.NullInt64
}

func InitStudent(user *userauth.User) *Student {
	student := &Student{
		User: user,
	}
	student.User.Type.Scan("student")

	return student
}
