package model

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id" bson:"id"`
	CreatedAt time.Time  `json:"-" bson:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"-" bson:"deleted_at,omitempty"`
}
