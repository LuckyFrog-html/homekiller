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
	Name     string
	Login    string
	Password string
	Subjects []*Subject `gorm:"many2many:teacher_to_subjects;"`
	Groups   []Group
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
	Description     string
	LessonID        uint
	Deadline        time.Time
	MaxScore        int
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
	Text                string
	HomeworkAnswerFiles []HomeworkAnswerFile
	TeacherResume       []TeacherResume
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
	TeacherResumeFiles []TeacherResumeFile
}

type TeacherResumeFile struct {
	gorm.Model
	TeacherResumeID uint
	Filepath        string
}
