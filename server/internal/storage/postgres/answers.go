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

func (s *Storage) GetHomeworkSolutions(homeworkId uint) ([]models.HomeworkAnswer, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var solutions []models.HomeworkAnswer

	result := tx.Preload("Student").Preload("HomeworkAnswerFiles").Preload("Homework").Find(&solutions, "homework_id = ?", homeworkId)

	return solutions, result.Error
}

func (s *Storage) AddHomeworkSolveReview(solveId uint, comment string, score int, teacherId uint) (models.TeacherResume, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	teacherResume := models.TeacherResume{HomeworkAnswerID: solveId, Comment: comment, Score: score, TeacherID: teacherId}
	result := tx.Create(&teacherResume)
	return teacherResume, result.Error
}
