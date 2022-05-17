package studentauth

import "github.com/jmoiron/sqlx"

func AddStudent(db *sqlx.DB, student *Student) {
	tx := db.MustBegin()
	tx.MustExec(
		`UPDATE "user"
		SET type = $1
		WHERE username = $2;`,
		student.User.Type,
		student.User.Username,
	)
	tx.Commit()
}
