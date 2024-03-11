package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	Name     string
	Stage    int64
	Login    string
	Password string
	Groups   []*Group `gorm:"many2many:students_to_groups;"`
}

type Teacher struct {
	gorm.Model
	Name    string
	Subject []*Subject `gorm:"many2many:teacher_to_subjects;"`
	Groups  []Group
}

type Group struct {
	gorm.Model
	Title     string
	IsActive  bool
	TeacherID uint
	Students  []*Student `gorm:"many2many:students_to_groups;"`
	Lessons   []Lesson
}

type Subject struct {
	gorm.Model
	Title   string
	Teacher []*Teacher `gorm:"many2many:teacher_to_subjects;"`
}

type Lesson struct {
	gorm.Model
	Date      time.Time
	GroupID   uint
	Homeworks []Homework
}

type Homework struct {
	gorm.Model
	description     string
	LessonID        uint
	deadline        time.Time
	MaxScore        string
	HomeworkFiles   []HomeworkFile
	HomeworkAnswers []HomeworkAnswer
}

type HomeworkFile struct {
	gorm.Model
	HomeworkID uint
	Filepath   string
}

type HomeworkAnswer struct {
	gorm.Model
	HomeworkID          uint
	text                string
	HomeworkAnswerFiles []HomeworkAnswerFile
	TeacherResume       []TeacherResume
}

type HomeworkAnswerFile struct {
	gorm.Model
	HomeworkAnswerID uint
	filepath         string
}

type TeacherResume struct {
	gorm.Model
	HomeworkAnswerID   uint
	TeacherResumeFiles []TeacherResumeFile
}

type TeacherResumeFile struct {
	gorm.Model
	TeacherResumeID uint
	filepath        string
}
