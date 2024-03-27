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
