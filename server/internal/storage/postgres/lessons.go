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

func (s *Storage) GetLessonById(lessonId uint) (*models.Lesson, error) {
	var lesson models.Lesson
	result := s.Db.Preload("Homeworks").Preload("Students").Preload("Groups").
		Raw("SELECT * FROM lessons WHERE id=? LIMIT 1;", lessonId).Scan(&lesson)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lesson, nil
}
