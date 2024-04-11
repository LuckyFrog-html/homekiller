package postgres

import (
	"fmt"
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
