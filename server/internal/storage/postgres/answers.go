package postgres

import (
	"fmt"
	"server/models"
)

func (s *Storage) GetSolutionById(id uint) (models.HomeworkAnswer, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var solution models.HomeworkAnswer

	result := tx.Preload("Homework").Preload("Student").First(&solution, "id = ?", id)

	return solution, result.Error
}

func (s *Storage) AddSolutionFile(solutionId uint, extension string) (uint, error) {
	tx := s.Db.Begin()
	solutionFile := models.HomeworkAnswerFile{HomeworkAnswerID: solutionId, Filepath: ""}
	result := tx.Create(&solutionFile)
	solutionFile.Filepath = fmt.Sprintf("files/students/%d.%s", solutionFile.ID, extension)
	tx.Save(&solutionFile)
	tx.Commit()

	return solutionFile.ID, result.Error
}