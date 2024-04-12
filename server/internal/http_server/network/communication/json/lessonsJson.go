package json

import "time"

type AddLessonJson struct {
	Date time.Time `json:"date"`
}

type MarkStudentAttendanceJson struct {
	StudentIDs []uint `json:"students_ids"`
}
