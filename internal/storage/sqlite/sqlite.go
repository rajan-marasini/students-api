package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rajan-marasini/students-api/internal/config"
	"github.com/rajan-marasini/students-api/internal/types"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER,
		email TEXT
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, email, age)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *Sqlite) GetStudentByID(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students WHERE id=? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Age, &student.Email)
	if err != nil {

		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}

		return types.Student{}, nil
	}

	return student, nil
}

func (s *Sqlite) GetAllStudents() ([]types.Student, error) {

	stmt, err := s.Db.Prepare("SELECT id, name, age, email FROM students")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Email)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (s *Sqlite) DeleteAStudent(id int64) (int64, error) {
	stmt, err := s.Db.Prepare("DELETE FROM students WHERE ID=?")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		return -1, err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	if rowAffected == 0 {
		return -1, fmt.Errorf("no student found with id:%d", id)
	}
	return rowAffected, nil
}

func (s *Sqlite) UpdateAStudent(student types.Student) (types.Student, error) {
	stmt, err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(&student.Name, &student.Email, &student.Age, &student.Id)
	if err != nil {
		return types.Student{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return types.Student{}, err
	}

	if rowsAffected == 0 {
		return types.Student{}, fmt.Errorf("no student found with id %d", &student.Id)
	}

	// Return the updated student
	student = types.Student{
		Id:    student.Id,
		Name:  student.Name,
		Email: student.Email,
		Age:   student.Age,
	}

	return student, nil
}
