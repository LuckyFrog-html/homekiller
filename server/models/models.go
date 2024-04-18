package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	Name             string
	Stage            int64
	Login            string `gorm:"unique;"`
	Password         string
	Lessons          []*Lesson `gorm:"many2many:students_to_lessons;"`
	HomeworksAnswers []*HomeworkAnswer
	Groups           []*Group `gorm:"many2many:students_to_groups;"`
}

func (s *Student) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(s.Password), []byte(password))
	return err == nil
}

type Teacher struct {
	gorm.Model
	Name           string
	Login          string `gorm:"unique;"`
	Password       string
	TeacherResumes []*TeacherResume
	Subjects       []*Subject `gorm:"many2many:teacher_to_subjects;"`
	Groups         []*Group   `gorm:"foreignKey:TeacherID;"`
}

func (t *Teacher) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(t.Password), []byte(password)) == nil
}

type Group struct {
	gorm.Model
	Title     string
	IsActive  bool
	TeacherID uint
	Students  []*Student `gorm:"many2many:students_to_groups;"`
	Lessons   []*Lesson  `gorm:"foreignKey:GroupID;"`
	Teacher   *Teacher
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
	Students  []*Student `gorm:"many2many:students_to_lessons;"`
	Group     *Group
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
	Lesson          *Lesson
}

type HomeworkFile struct {
	gorm.Model
	HomeworkID uint
	Filepath   string
	Homework   *Homework
}

type HomeworkAnswer struct {
	gorm.Model
	Text                string
	HomeworkID          uint
	StudentID           uint
	HomeworkAnswerFiles []*HomeworkAnswerFile
	TeacherResumes      []*TeacherResume
	Student             *Student
	Homework            *Homework
}

type HomeworkAnswerFile struct {
	gorm.Model
	HomeworkAnswerID uint
	Filepath         string
	HomeworkAnswer   *HomeworkAnswer
}

type TeacherResume struct {
	gorm.Model
	HomeworkAnswerID   uint
	Comment            string
	Score              int
	TeacherID          uint
	TeacherResumeFiles []*TeacherResumeFile
	Teacher            *Teacher
	HomeworkAnswer     *HomeworkAnswer
}

type TeacherResumeFile struct {
	gorm.Model
	TeacherResumeID uint
	Filepath        string
	TeacherResume   *TeacherResume
}

type StudentsToLessons struct {
	gorm.Model
	StudentID uint
	LessonID  uint
	Group     *Group
	Lesson    *Lesson
}
