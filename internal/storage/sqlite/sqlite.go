package sqlite

import (
	"database/sql"

	"github.com/Ashank007/students-api-go/internal/config"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Ashank007/students-api-go/internal/types"

)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite,error) {

	db,err := sql.Open("sqlite3",cfg.StoragePath)

	if err != nil {
		return nil,err
	}

	_,err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)

	if err != nil {
		return nil,err
  }

	return &Sqlite{
    Db: db,
	},nil

}

func (s *Sqlite) CreateStudent(name string,email string,age int) (int64,error) {

	stmt,err := s.Db.Prepare("INSERT INTO students (name,email,age) values (?,?,?)")

	if err != nil{
		return 0,err
	}

	defer stmt.Close()

	result,err := stmt.Exec(name,email,age)

	if err != nil {
		return 0,nil
	}

	lastId,err := result.LastInsertId()

	if err != nil {
    return 0,err
	}

	return lastId,nil
}

func (s *Sqlite) GetAllStudents() ([]map[string]interface{}, error) {
	rows, err := s.Db.Query("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []map[string]interface{}

	for rows.Next() {
		var id int64
		var name, email string
		var age int
		if err := rows.Scan(&id, &name, &email, &age); err != nil {
			return nil, err
		}

		students = append(students, map[string]interface{}{
			"id":    id,
			"name":  name,
			"email": email,
			"age":   age,
		})
	}

	return students, nil
}

func (s *Sqlite) UpdateStudent(student types.Student) error {
	_, err := s.Db.Exec(`UPDATE students SET name=?, email=?, age=? WHERE id=?`, student.Name, student.Email, student.Age, student.Id)
	return err
}

func (s *Sqlite) DeleteStudent(id int64) error {
	_, err := s.Db.Exec(`DELETE FROM students WHERE id=?`, id)
	return err
}

