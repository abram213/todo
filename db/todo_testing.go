package db

import (
	"github.com/jinzhu/copier"
	"todo/model"
)

func (db *TestDatabase) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	var todoCopy model.Todo
	//test full struct copy
	if err := copier.Copy(&todoCopy, todo); err != nil {
		return nil, err
	}
	// todoCopy.ID = 7
	return &todoCopy, nil
}

func (db *TestDatabase) UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	return todo, nil
}

func (db *TestDatabase) GetTodo(id uint) (*model.Todo, error) {
	todo := model.Todo{Model: model.Model{ID: id}}
	return &todo, nil
}

func (db *TestDatabase) GetTodos() (*[]model.Todo, error) {
	return db.Todos, nil
}

func (db *TestDatabase) GetActiveTodos() (*[]model.Todo, error) {
	return db.Todos, nil
}

func (db *TestDatabase) GetDoneTodos() (*[]model.Todo, error) {
	return db.Todos, nil
}

func (db *TestDatabase) DeleteTodo(id uint) error {
	return nil
}
