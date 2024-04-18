package postgres

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server/models"
)

func (s *Storage) GetTeacher(login, password string) (models.Teacher, error) {
	tx := s.Db.Begin()
	var teacher models.Teacher

	result := tx.Preload("Subjects").Preload("Groups").First(&teacher, "login = ?", login)

	if result.Error != nil {
		return models.Teacher{}, result.Error
	}

	if !teacher.CheckPassword(password) {
		return models.Teacher{}, fmt.Errorf("password is incorrect")
	}

	return teacher, nil
}

func (s *Storage) GetTeacherById(id int) (models.Teacher, error) {
	tx := s.Db.Begin()
	var teacher models.Teacher

	result := tx.First(&teacher, id)

	if result.Error != nil {
		return models.Teacher{}, result.Error
	}

	return teacher, nil
}

func (s *Storage) AddTeacher(name, login, password string) models.Teacher {
	tx := s.Db.Begin()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	teacher := models.Teacher{
		Name:     name,
		Login:    login,
		Password: string(bytes),
	}

	tx.Create(&teacher)
	tx.Commit()

	return teacher
}
