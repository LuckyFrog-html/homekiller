package postgres

import (
	"server/models"
	"time"
)

func (s *Storage) AddGroup(title string, teacherId uint) models.Group {
	tx := s.Db.Begin()
	group := models.Group{Title: title, TeacherID: teacherId, IsActive: true}

	tx.Create(&group)
	tx.Commit()

	return group
}

func (s *Storage) GetGroupById(id uint) (models.Group, error) {
	tx := s.Db.Begin()
	var group models.Group

	result := tx.Preload("Students").First(&group, "id=?", id)
	if result.Error != nil {
		return models.Group{}, result.Error
	}

	return group, nil
}

func (s *Storage) AddStudentsToGroup(groupId uint, studentsIds []uint) {
	tx := s.Db.Begin()
	for _, studentId := range studentsIds {
		studentToGroup := models.StudentsToGroups{StudentID: studentId, GroupID: groupId, AppendDate: time.Now()}

		tx.Create(&studentToGroup)
	}
	tx.Commit()
}

func (s *Storage) IsStudentInGroup(groupId, studentId uint) bool {
	tx := s.Db.Begin()
	var studentToGroup models.StudentsToGroups

	result := tx.First(&studentToGroup, "group_id = ? AND student_id = ?", groupId, studentId)

	return result.Error == nil
}
