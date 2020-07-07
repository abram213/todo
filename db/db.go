package db

import "todo/model"

type DataLayer interface {
	CreateTodo(todo *model.Todo) (*model.Todo, error)
	UpdateTodo(todo *model.Todo) (*model.Todo, error)
	GetActiveTodos() (*[]model.Todo, error)
	GetDoneTodos() (*[]model.Todo, error)
	GetTodos() (*[]model.Todo, error)
	GetTodo(id uint) (*model.Todo, error)
	DeleteTodo(id uint) error

	Migrate(values ...interface{})
	DropTables(values ...interface{})
	CloseDB() error
}
