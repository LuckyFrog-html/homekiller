package json

import "time"

type HomeworkJson struct {
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	MaxScore    int       `json:"max_score"`
	LessonId    uint      `json:"lesson_id"`
}

type HomeworkAnswerJson struct {
	HomeworkId uint   `json:"homework_id"`
	Text       string `json:"text"`
}
