package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
	"todo/model"
)

func (db *FileDatabase) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	db.data.lastID++
	todo.ID = db.data.lastID
	todo.UpdatedAt = time.Now()

	db.data.todos[todo.ID] = todo
	return todo, nil
}

func (db *FileDatabase) UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	if _, ok := db.data.todos[todo.ID]; !ok {
		return nil, fmt.Errorf("cant find todo with id: %v", todo.ID)
	}
	todo.UpdatedAt = time.Now()
	db.data.todos[todo.ID] = todo
	return db.data.todos[todo.ID], nil
}

func (db *FileDatabase) GetTodo(id uint) (*model.Todo, error) {
	todo, ok := db.data.todos[id]
	if !ok {
		return nil, fmt.Errorf("cant find todo with id: %v", todo.ID)
	}
	return todo, nil
}

func (db *FileDatabase) GetTodos() (*[]model.Todo, error) {
	todos := make([]model.Todo, 0, len(db.data.todos))
	for _, todo := range db.data.todos {
		todos = append(todos, *todo)
	}
	sort.Slice(todos, func(i, j int) bool { return todos[i].UpdatedAt.Before(todos[j].UpdatedAt) })
	return &todos, nil
}

func (db *FileDatabase) GetActiveTodos() (*[]model.Todo, error) {
	var todos []model.Todo
	for _, todo := range db.data.todos {
		if !todo.Done {
			todos = append(todos, *todo)
		}
	}
	sort.Slice(todos, func(i, j int) bool { return todos[i].UpdatedAt.Before(todos[j].UpdatedAt) })
	return &todos, nil
}

func (db *FileDatabase) GetDoneTodos() (*[]model.Todo, error) {
	var todos []model.Todo
	for _, todo := range db.data.todos {
		if todo.Done {
			todos = append(todos, *todo)
		}
	}
	sort.Slice(todos, func(i, j int) bool { return todos[i].ID < todos[j].ID })
	return &todos, nil
}

func (db *FileDatabase) DeleteTodo(id uint) error {
	delete(db.data.todos, id)
	return nil
}

func (db *FileDatabase) saveTodos() error {
	todos := make([]model.Todo, 0, len(db.data.todos))
	for _, todo := range db.data.todos {
		todos = append(todos, *todo)
	}
	sort.Slice(todos, func(i, j int) bool { return todos[i].UpdatedAt.Before(todos[j].UpdatedAt) })
	var data string
	for _, todo := range todos {
		data += todo.String() + "\n"
	}

	err := ioutil.WriteFile(db.path, []byte(data), os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}
