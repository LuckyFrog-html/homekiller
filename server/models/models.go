package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	Name             string
	Stage            int64
	Login            string
	Password         string
	Lessons          []*Lesson `gorm:"many2many:students_to_lessons;"`
	HomeworksAnswers []*HomeworkAnswer
	Groups           []*Group `gorm:"many2many:students_to_groups;"`
}

type Teacher struct {
	gorm.Model
	Name           string
	Login          string
	Password       string
	TeacherResumes []*TeacherResume
	Subjects       []*Subject `gorm:"many2many:teacher_to_subjects;"`
	Groups         []*Group
}

type Group struct {
	gorm.Model
	Title     string
	IsActive  bool
	TeacherID uint
	Students  []*Student `gorm:"many2many:students_to_groups;"`
	Lessons   []*Lesson
}

type Subject struct {
	gorm.Model
	Title    string
	Teachers []*Teacher `gorm:"many2many:teacher_to_subjects;"`
}

type Lesson struct {
	gorm.Model
	Date      time.Time
	GroupID   uint
	Homeworks []*Homework
	Student   []*Student `gorm:"many2many:students_to_lessons;"`
}

type StudentsToGroups struct {
	StudentID  uint
	GroupID    uint
	AppendDate time.Time
}

type Homework struct {
	gorm.Model
	Description     string
	LessonID        uint
	Deadline        time.Time
	MaxScore        int
	HomeworkFiles   []*HomeworkFile
	HomeworkAnswers []*HomeworkAnswer
}

type HomeworkFile struct {
	gorm.Model
	HomeworkID uint
	Filepath   string
}

type HomeworkAnswer struct {
	gorm.Model
	Text                string
	HomeworkID          uint
	StudentID           uint
	HomeworkAnswerFiles []*HomeworkAnswerFile
	TeacherResumes      []*TeacherResume
}

type HomeworkAnswerFile struct {
	gorm.Model
	HomeworkAnswerID uint
	Filepath         string
}

type TeacherResume struct {
	gorm.Model
	HomeworkAnswerID   uint
	Comment            string
	Score              int
	TeacherID          uint
	TeacherResumeFiles []*TeacherResumeFile
}

type TeacherResumeFile struct {
	gorm.Model
	TeacherResumeID uint
	Filepath        string
}
