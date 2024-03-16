package services

import (
	"fmt"
	"server/internal/storage/postgres"
	"server/models"
)

func GetAll(s *postgres.Storage, id string) error {
	const op = "storage.postgres.GetAll"

	var student models.Student
	var _ []models.Student

	_ = s.Db.Table("students").First(&student)

	return nil
}

func AddStudent(s *postgres.Storage, name string, stage int64, login, password string) (int64, error) {
	res := s.Db.Create(&models.Student{Name: name, Stage: stage, Login: login, Password: password})
	if res.Error != nil {
		return -1, res.Error
	}
	fmt.Println(res.RowsAffected)
	return 0, nil
}
