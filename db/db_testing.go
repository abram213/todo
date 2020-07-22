package db

import (
	"todo/model"
)

type TestDatabase struct {
	Todo  model.Todo
	Todos *[]model.Todo
}

func (db *TestDatabase) CloseDB() error                   { return nil }
func (db *TestDatabase) Migrate(values ...interface{})    {}
func (db *TestDatabase) DropTables(values ...interface{}) {}
