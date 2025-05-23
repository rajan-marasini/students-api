package storage

import "github.com/rajan-marasini/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentByID(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
	DeleteAStudent(id int64) (int64, error)
	UpdateAStudent(student types.Student) (types.Student, error)
}
