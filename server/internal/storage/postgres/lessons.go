package postgres

import (
	"server/models"
	"time"
)

func (s *Storage) AddLesson(date time.Time, groupId uint) *models.Lesson {
	lesson := &models.Lesson{
		Date:    date,
		GroupID: groupId,
	}
	s.Db.Create(lesson)
	s.Db.Commit()
	return lesson
}
