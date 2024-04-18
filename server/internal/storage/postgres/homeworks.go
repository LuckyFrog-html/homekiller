package postgres

import (
	"server/models"
	"time"
)

func (s *Storage) AddHomework(description string, lessonId uint, deadline time.Time, maxScore int) models.Homework {
	homework := models.Homework{Description: description, LessonID: lessonId, Deadline: deadline, MaxScore: maxScore}

	s.Db.Create(&homework)
	s.Db.Commit()

	return homework
}
