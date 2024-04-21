package postgres

import (
	"server/models"
	"time"
)

func (s *Storage) AddHomework(description string, lessonId uint, deadline time.Time, maxScore int) (models.Homework, error) {
	tx := s.Db.Begin()
	homework := models.Homework{Description: description, LessonID: lessonId, Deadline: deadline, MaxScore: maxScore}

	result := tx.Create(&homework)
	tx.Commit()

	return homework, result.Error
}

func (s *Storage) IsHomeworkInLesson(homeworkId, lessonId uint) bool {
	var homework models.Homework
	tx := s.Db.Begin()
	result := tx.First(&homework, "id = ? AND lesson_id = ?", homeworkId, lessonId)
	tx.Commit()
	return result.RowsAffected != 0
}

func (s *Storage) AddHomeworkFiles(homeworkId uint, files []string) error {
	tx := s.Db.Begin()
	for _, file := range files {
		homeworkFile := models.HomeworkFile{HomeworkID: homeworkId, Filepath: file}
		tx.Create(&homeworkFile)
	}
	tx.Commit()
	return tx.Error
}

type StudentHomework struct {
	models.Homework
	IsDone     bool
	GroupId    int
	GroupTitle string
}

func (s *Storage) GetHomeworksByStudent(studentId uint) ([]StudentHomework, error) {
	tx := s.Db.Begin()
	var studentsHomeworks []StudentHomework
	result := tx.Raw("SELECT hw.*, ls.group_id, ha.id IS NOT NULL AS IsDone, g2.title AS group_title FROM homeworks hw JOIN lessons ls ON hw.lesson_id = ls.id JOIN students_to_groups g on g.group_id = ls.group_id LEFT OUTER JOIN homework_answers ha ON ha.homework_id = hw.id AND ha.student_id = g.student_id JOIN public.groups g2 ON ls.group_id = g2.id WHERE g.student_id=?", studentId).Scan(&studentsHomeworks)
	return studentsHomeworks, result.Error
}
