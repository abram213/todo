package model

import (
	"fmt"
	"time"
)

type Todo struct {
	Model
	Body     string    `json:"body" gorm:"type:text" bson:"body"`
	Done     bool      `json:"done" gorm:"default:false" bson:"done"`
	DoneTime time.Time `json:"done_time" bson:"done_time"`
}

func (t Todo) String() string {
	updateTime := t.UpdatedAt.Format("02-01-2006/15:04:05")
	doneTime := t.DoneTime.Format("02-01-2006/15:04:05")
	return fmt.Sprintf("%v | %v | %v | %v | %v", t.ID, updateTime, t.Body, t.Done, doneTime)
}
