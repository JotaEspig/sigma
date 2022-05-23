package studentauth

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Adds a student to a database.
// Panics if something goes wrong.
func AddStudent(db *sqlx.DB, s *Student) {
	tx := db.MustBegin()
	tx.MustExec(
		`UPDATE "user"
		SET type = 'student'
		WHERE id = $1;`,
		s.ID,
	)
	tx.MustExec(
		`INSERT INTO student(id, year, status, class_id)
		VALUES($1, $2, $3, $4)`,
		s.ID, s.Year, s.Status, s.ClassID,
	)
	tx.Commit()
}

// Gets a student from a database
func GetStudent(db *sqlx.DB, id int, columns ...string) (*Student, error) {
	s := Student{}

	columnsStr := strings.Join(columns, ",")
	// If there is no arguments, sets it to default
	if len(columns) == 0 {
		columnsStr = "*"
	}

	// TODO Jota: Implement something to "clean" columnsStr
	// the reason is to avoid SQLInjection (if that's possible)
	sqlQuery := fmt.Sprintf("SELECT %s FROM \"student\" WHERE id=$1", columnsStr)

	err := db.Get(&s, sqlQuery, id)

	return &s, err
}
