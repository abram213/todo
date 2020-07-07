package app

import (
	"time"
	"todo/model"
)

func (app *App) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	return app.Database.CreateTodo(todo)
}

func (app *App) UpdateTodo(id uint, body string) (*model.Todo, error) {
	todo, err := app.GetTodo(id)
	if err != nil {
		return nil, err
	}
	todo.Body = body
	return app.Database.UpdateTodo(todo)
}

func (app *App) GetTodo(id uint) (*model.Todo, error) {
	return app.Database.GetTodo(id)
}

func (app *App) GetActiveTodos() (*[]model.Todo, error) {
	return app.Database.GetActiveTodos()
}

func (app *App) GetDoneTodos() (*[]model.Todo, error) {
	return app.Database.GetDoneTodos()
}

func (app *App) GetTodos() (*[]model.Todo, error) {
	return app.Database.GetTodos()
}

func (app *App) DeleteTodo(id uint) error {
	return app.Database.DeleteTodo(id)
}

func (app *App) DoneTodo(id uint) (*model.Todo, error) {
	todo, err := app.GetTodo(id)
	if err != nil {
		return nil, err
	}
	todo.Done = true
	todo.DoneTime = time.Now()
	return app.Database.UpdateTodo(todo)
}

func (app *App) UndoneTodo(id uint) (*model.Todo, error) {
	todo, err := app.GetTodo(id)
	if err != nil {
		return nil, err
	}
	todo.Done = false
	return app.Database.UpdateTodo(todo)
}
