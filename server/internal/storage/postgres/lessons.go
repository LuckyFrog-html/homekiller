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

func (s *Storage) GetLessonsByStudentId(studentId uint) ([]*models.Lesson, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var student models.Student
	result := tx.Preload("Lessons").First(&student, studentId)

	return student.Lessons, result.Error
}

func (s *Storage) IsStudentInLesson(lessonId, studentId uint) (bool, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var lesson models.Lesson
	result := tx.Preload("Students").
		Joins("JOIN students_to_groups sg ON sg.group_id = lessons.group_id").
		First(&lesson, "sg.student_id = ? AND lessons.lesson_id = ?", studentId, lessonId)
	return result.RowsAffected != 0, result.Error
}

func (s *Storage) GetLessonsByTeacher(teacherId uint) ([]*models.Lesson, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var lessons []*models.Lesson
	result := tx.Preload("Group").Preload("Homeworks").Preload("Students").
		Joins("JOIN groups ON groups.id = lessons.group_id").
		Where("groups.teacher_id = ?", teacherId).Find(&lessons)
	return lessons, result.Error
}
