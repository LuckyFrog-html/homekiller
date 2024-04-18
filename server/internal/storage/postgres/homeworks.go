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
