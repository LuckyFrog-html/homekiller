package postgres

import (
	"server/models"
	"time"
)

func (s *Storage) AddGroup(title string, teacherId uint) models.Group {
	group := models.Group{Title: title, TeacherID: teacherId, IsActive: true}

	s.Db.Create(&group)
	s.Db.Commit()

	return group
}

func (s *Storage) GetGroupById(id uint) (models.Group, error) {
	var group models.Group

	result := s.Db.First(&group, id)

	if result.Error != nil {
		return models.Group{}, result.Error
	}

	return group, nil
}

func (s *Storage) AddStudentsToGroup(groupId uint, studentsIds []uint) {
	for _, studentId := range studentsIds {
		studentToGroup := models.StudentsToGroups{StudentID: studentId, GroupID: groupId, AppendDate: time.Now()}

		s.Db.Create(&studentToGroup)
	}
	s.Db.Commit()
}
