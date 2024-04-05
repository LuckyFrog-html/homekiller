package postgres

import "server/models"

func (s *Storage) GetAllStudents() error {
	const op = "storage.postgres.GetAllStudents"

	var student models.Student

	students := s.Db.Table("students").First(&student)

	print(students)

	return nil
}

func (s *Storage) GetStudentByID(id uint) error {
	const op = "storage.postgres.GetStudentByID"

	var student models.Student

	students := s.Db.Table("students").First(&student, id)

	print(students)

	return nil
}

func (s *Storage) AddStudent(name string, stage int64, login, password string) models.Student {
	const op = "storage.postgres.AddStudent"

	student := models.Student{
		Name:     name,
		Stage:    stage,
		Login:    login,
		Password: password,
	}

	result := s.Db.Create(&student)

	print(result)

	return student
}
