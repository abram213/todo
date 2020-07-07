package db

import (
	"github.com/pkg/errors"
	"todo/model"
)

func (db *SQLDatabase) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	return todo, errors.Wrap(db.Create(&todo).Error, "unable to create todo")
}

func (db *SQLDatabase) UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	return todo, errors.Wrap(db.Save(&todo).Error, "unable to update todo")
}

func (db *SQLDatabase) GetTodo(id uint) (*model.Todo, error) {
	var todo model.Todo
	return &todo, errors.Wrap(db.First(&todo, id).Error, "unable to get todo")
}

func (db *SQLDatabase) GetTodos() (*[]model.Todo, error) {
	var todos []model.Todo
	return &todos, errors.Wrap(db.Find(&todos).Error, "unable to get todos")
}

func (db *SQLDatabase) GetActiveTodos() (*[]model.Todo, error) {
	var todos []model.Todo
	return &todos, errors.Wrap(db.Where("done = ?", false).Find(&todos).Error, "unable to get active todos")
}

func (db *SQLDatabase) GetDoneTodos() (*[]model.Todo, error) {
	var todos []model.Todo
	return &todos, errors.Wrap(db.Where("done = ?", true).Find(&todos).Error, "unable to get done todos")
}

func (db *SQLDatabase) DeleteTodo(id uint) error {
	todo := model.Todo{Model: model.Model{ID: id}}
	return errors.Wrap(db.Delete(&todo).Error, "unable to delete todo")
}
