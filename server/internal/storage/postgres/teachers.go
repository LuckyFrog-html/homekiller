package postgres

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server/models"
)

func (s *Storage) GetTeacher(login, password string) (models.Teacher, error) {
	const op = "storage.postgres.GetStudentByLogin"

	var teacher models.Teacher

	result := s.Db.First(&teacher, "login = ?", login)

	if result.Error != nil {
		return models.Teacher{}, result.Error
	}

	if !teacher.CheckPassword(password) {
		return models.Teacher{}, fmt.Errorf("password is incorrect")
	}

	return teacher, nil
}

func (s *Storage) GetTeacherById(id int) (models.Teacher, error) {
	var teacher models.Teacher

	result := s.Db.First(&teacher, id)

	if result.Error != nil {
		return models.Teacher{}, result.Error
	}

	return teacher, nil
}

func (s *Storage) AddTeacher(name, login, password string) models.Teacher {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	teacher := models.Teacher{
		Name:     name,
		Login:    login,
		Password: string(bytes),
	}

	s.Db.Create(&teacher)
	s.Db.Commit()

	return teacher
}
