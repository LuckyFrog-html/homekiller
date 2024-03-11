package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"server/models"
)

func New(dbString string) (*gorm.DB, error) {
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

	return db, nil
}
