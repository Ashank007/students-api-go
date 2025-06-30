package storage

import "github.com/Ashank007/students-api-go/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetAllStudents() ([]map[string]interface{}, error)
	UpdateStudent(student types.Student) error
	DeleteStudent(id int64) error
}


