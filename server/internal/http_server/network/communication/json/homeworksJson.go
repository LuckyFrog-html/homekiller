package json

import "time"

type HomeworkJson struct {
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	MaxScore    int       `json:"max_score"`
}
