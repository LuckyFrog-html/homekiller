package postgres

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server/models"
)

func (s *Storage) GetAllStudents() error {
	const op = "storage.postgres.GetAllStudents"

	var student models.Student

	students := s.Db.Table("students").First(&student)

	print(students)

	return nil
}

func (s *Storage) GetStudentByID(id uint) (*models.Student, error) {
	const op = "storage.postgres.GetStudentByID"

	var student models.Student

	result := s.Db.First(&student, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}

func (s *Storage) AddStudent(name string, stage int64, login, password string) models.Student {
	const op = "storage.postgres.AddStudent"

	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	student := models.Student{
		Name:     name,
		Stage:    stage,
		Login:    login,
		Password: string(bytes),
	}

	s.Db.Create(&student)
	s.Db.Commit()

	return student
}

func (s *Storage) GetStudentByLogin(login, password string) (models.Student, error) {
	const op = "storage.postgres.GetStudentByLogin"

	var student models.Student

	result := s.Db.Preload("Lessons").Preload("Groups").Preload("HomeworksAnswers").First(&student, "login = ?", login)

	if result.Error != nil {
		return models.Student{}, result.Error
	}

	if !student.CheckPassword(password) {
		return models.Student{}, fmt.Errorf("password is incorrect")
	}

	return student, nil
}

func (s *Storage) GetStudentsByGroup(groupId uint) ([]*models.Student, error) {
	var group models.Group

	result := s.Db.Preload("Students").First(&group, groupId)
	if result.Error != nil {
		return nil, result.Error
	}

	return group.Students, nil
}

func (s *Storage) MarkStudentAttendance(studentsIds []uint, lessonId uint) {
	for _, studentId := range studentsIds {
		studentToLesson := models.StudentsToLessons{StudentID: studentId, LessonID: lessonId}

		s.Db.Create(&studentToLesson)
	}
	s.Db.Commit()
}
