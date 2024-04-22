package json

import "time"

type AddLessonJson struct {
	Date    time.Time `json:"date"`
	GroupId uint      `json:"group_id"`
}

type MarkStudentAttendanceJson struct {
	StudentIDs []uint `json:"students_ids"`
	LessonID   uint   `json:"lesson_id"`
}
