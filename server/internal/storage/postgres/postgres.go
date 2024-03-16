package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"server/models"
)

type Storage struct {
	Db *gorm.DB
}

func New(dbString string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.AutoMigrate(&models.Student{}, &models.Teacher{}, &models.Group{},
		&models.Subject{}, &models.Lesson{}, &models.Homework{},
		&models.HomeworkAnswer{}, &models.HomeworkFile{},
		&models.HomeworkAnswerFile{}, &models.TeacherResume{}, &models.TeacherResumeFile{})

	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, nil
}

func (s *Storage) GetAllStudents(id string) error {
	const op = "storage.postgres.GetAll"

	var student models.Student
	var _ []models.Student

	_ = s.Db.Table("students").First(&student)

	return nil
}
