package services

import (
	"server/internal/storage/postgres"
	"server/models"
)

func (s *postgres.Storage) GetAll(id string) error {
	const op = "storage.postgres.GetAll"

	var student models.Student
	var _ []models.Student

	_ = s.db.Table("students").First(&student)

	return nil
}
