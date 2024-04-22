package postgres

import (
	"server/models"
	"time"
)

func (s *Storage) AddLesson(date time.Time, groupId uint) (*models.Lesson, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	lesson := models.Lesson{
		Date:    date,
		GroupID: groupId,
	}
	result := tx.Create(&lesson)
	tx.Commit()
	return &lesson, result.Error
}

func (s *Storage) GetLessonById(lessonId uint) (*models.Lesson, error) {
	var lesson models.Lesson

	tx := s.Db.Begin()
	defer tx.Commit()
	result := tx.Preload("Group").Preload("Homeworks").Preload("Students").First(&lesson, "id=?", lessonId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lesson, nil
}

func (s *Storage) GetLessonByHomeworkId(homeworkId uint) (*models.Lesson, error) {
	var lesson models.Lesson

	tx := s.Db.Begin()
	defer tx.Commit()
	result := tx.Preload("Group").Preload("Homeworks").Preload("Students").Joins("JOIN homeworks ON homeworks.lesson_id = lessons.id").First(&lesson, "homeworks.id=?", homeworkId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lesson, nil
}
