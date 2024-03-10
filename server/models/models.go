package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string // A regular string field
	Stage    int64  // A pointer to a string, allowing for null values
	Login    string // An unsigned 8-bit integer
	Password string // A pointer to time.Time, can be null
}
