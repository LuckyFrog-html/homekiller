package postgres

import (
	"server/models"
	"time"
)

func (s *Storage) AddHomework(description string, lessonId uint, deadline time.Time, maxScore int) models.Homework {
	tx := s.Db.Begin()
	homework := models.Homework{Description: description, LessonID: lessonId, Deadline: deadline, MaxScore: maxScore}

	tx.Create(&homework)
	tx.Commit()

	return homework
}

func (s *Storage) IsHomeworkInLesson(homeworkId, lessonId uint) bool {
	var homework models.Homework
	tx := s.Db.Begin()
	result := tx.First(&homework, "id = ? AND lesson_id = ?", homeworkId, lessonId)
	tx.Commit()
	return result.RowsAffected != 0
}

func (s *Storage) AddHomeworkFiles(homeworkId uint, files []string) {
	tx := s.Db.Begin()
	for _, file := range files {
		homeworkFile := models.HomeworkFile{HomeworkID: homeworkId, Filepath: file}
		tx.Create(&homeworkFile)
	}
	tx.Commit()
}
