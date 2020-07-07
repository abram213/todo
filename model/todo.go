package model

import (
	"time"
)

type Todo struct {
	Model
	Body     string    `json:"body" gorm:"type:text" bson:"body"`
	Done     bool      `json:"done" gorm:"default:false" bson:"done"`
	DoneTime time.Time `json:"done_time" bson:"done_time"`
}
