package postgres

import (
	"fmt"
	"server/models"
	"strings"
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

func (s *Storage) AddHomeworkFile(homeworkId uint, file string) (uint, error) {
	tx := s.Db.Begin()
	homeworkFile := models.HomeworkFile{HomeworkID: homeworkId, Filepath: file}
	tx.Create(&homeworkFile)
	tmp := strings.Split(file, ".")
	homeworkFile.Filepath = fmt.Sprintf("files/%d.%s", homeworkFile.ID, tmp[len(tmp)-1])
	tx.Save(&homeworkFile)
	tx.Commit()
	return homeworkFile.ID, tx.Error
}

type StudentHomework struct {
	models.Homework
	IsDone       bool
	GroupId      int
	GroupTitle   string
	SubjectTitle string
}

func (s *Storage) GetHomeworksByStudent(studentId uint) ([]StudentHomework, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var studentsHomeworks []StudentHomework
	result := tx.Raw("SELECT hw.*, ls.group_id, ha.id IS NOT NULL AS IsDone, g2.title AS group_title, CASE WHEN s.title IS NULL THEN 'default' END AS subject_title FROM homeworks hw JOIN lessons ls ON hw.lesson_id = ls.id JOIN students_to_groups g on g.group_id = ls.group_id LEFT OUTER JOIN homework_answers ha ON ha.homework_id = hw.id AND ha.student_id = g.student_id JOIN public.groups g2 ON ls.group_id = g2.id JOIN teachers t ON t.id = g2.teacher_id LEFT OUTER JOIN teacher_to_subjects ts ON ts.teacher_id = t.id LEFT OUTER JOIN subjects s ON ts.subject_id = s.id WHERE g.student_id = ? ORDER BY hw.deadline, ha.id IS NOT NULL;", studentId).Scan(&studentsHomeworks)
	return studentsHomeworks, result.Error
}

func (s *Storage) GetHomeworkById(id uint) (models.Homework, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var homework models.Homework
	result := tx.Preload("Lesson").Preload("HomeworkFiles").First(&homework, id)
	return homework, result.Error
}

func (s *Storage) AddHomeworkAnswer(homeworkId, studentId uint, answer string) error {
	tx := s.Db.Begin()
	defer tx.Commit()
	homeworkAnswer := models.HomeworkAnswer{HomeworkID: homeworkId, StudentID: studentId, Text: answer}
	tx.Create(&homeworkAnswer)
	return tx.Error
}

func (s *Storage) GetHomeworkSolvesByTeacher(teacherId uint) ([]models.HomeworkAnswer, error) {
	tx := s.Db.Begin()
	defer tx.Commit()
	var homeworkAnswers []models.HomeworkAnswer
	result := tx.Preload("Homework").Preload("Student").Preload("HomeworkAnswerFiles").
		Joins("JOIN homeworks ON homeworks.id = homework_answers.homework_id").
		Joins("JOIN lessons ON lessons.id = homeworks.lesson_id").
		Joins("JOIN groups ON groups.id = lessons.group_id").
		Where("groups.teacher_id = ?", teacherId).Find(&homeworkAnswers)
	return homeworkAnswers, result.Error
}
