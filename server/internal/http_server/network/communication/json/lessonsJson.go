package json

import "time"

type AddLessonJson struct {
	Date time.Time `json:"date"`
}
