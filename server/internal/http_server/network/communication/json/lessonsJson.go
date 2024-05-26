package json

type AddLessonJson struct {
	Date    string `json:"date"`
	GroupId uint   `json:"group_id"`
}

type MarkStudentAttendanceJson struct {
	StudentIDs []uint `json:"students_ids"`
	LessonID   uint   `json:"lesson_id"`
}

type DeleteLessonJson struct {
	LessonID uint `json:"lesson_id"`
}
