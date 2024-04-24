package postgres

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server/models"
)

func (s *Storage) GetAllStudents() ([]models.Student, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var students []models.Student

	res := tx.Table("students").Find(&students)

	return students, res.Error
}

func (s *Storage) GetStudentByID(id uint) (*models.Student, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var student models.Student

	result := tx.Raw("SELECT * FROM students WHERE id = ? LIMIT 1;", id).Scan(&student)

	return &student, result.Error
}

func (s *Storage) AddStudent(name string, stage int64, login, password string) (*models.Student, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	tx := s.Db.Begin()
	defer tx.Commit()

	student := models.Student{
		Name:     name,
		Stage:    stage,
		Login:    login,
		Password: string(bytes),
	}

	result := tx.Create(&student)
	tx.Commit()

	return &student, result.Error
}

func (s *Storage) GetStudentByLogin(login, password string) (models.Student, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var student models.Student
	tmp := tx.Preload("Lessons").Preload("Groups").Preload("HomeworksAnswers")
	//result := tmp.Raw("SELECT * FROM students WHERE login = ? LIMIT 1;", login).Scan(&student)
	result := tmp.First(&student, "login = ?", login)

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
	tx := s.Db.Begin()
	defer tx.Commit()

	result := tx.Model(&group).Preload("Students").First(&group, groupId)
	if result.Error != nil {
		return nil, result.Error
	}

	return group.Students, nil
}

func (s *Storage) MarkStudentAttendance(studentsIds []uint, lessonId uint) {
	tx := s.Db.Begin()
	defer tx.Commit()
	for _, studentId := range studentsIds {
		studentToLesson := models.StudentsToLessons{StudentID: studentId, LessonID: lessonId}

		tx.Create(&studentToLesson)
	}
	tx.Commit()
}
