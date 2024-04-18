package postgres

import (
	"server/models"
	"time"
)

func (s *Storage) AddLesson(date time.Time, groupId uint) *models.Lesson {
	tx := s.Db.Begin()
	lesson := models.Lesson{
		Date:    date,
		GroupID: groupId,
	}
	tx.Create(&lesson)
	tx.Commit()
	return &lesson
}

func (s *Storage) GetLessonById(lessonId uint) (*models.Lesson, error) {
	var lesson models.Lesson

	tx := s.Db.Begin()
	result := tx.Preload("Group").Preload("Homeworks").Preload("Students").First(&lesson, "id=?", lessonId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lesson, nil
}
