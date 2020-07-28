package app

import (
	"reflect"
	"testing"
	"time"
	"todo/db"
	"todo/model"
)

func TestCreateTodo(t *testing.T) {
	tApp := App{Database: &db.TestDatabase{}}
	tTodo := &model.Todo{
		Model:    model.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Body:     "Test",
		Done:     true,
		DoneTime: time.Now(),
	}
	todo, err := tApp.CreateTodo(tTodo)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(tTodo, todo) {
		t.Errorf("\nUnexpected result got:\n %+v \nwant:\n %+v", todo, tTodo)
	}
}

func TestUpdateTodo(t *testing.T) {
	tApp := App{Database: &db.TestDatabase{}}
	id, body := 1, "TestBody"
	todo, err := tApp.UpdateTodo(uint(id), body)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if todo.ID != uint(id) {
		t.Errorf("\nUnexpected id got: %#v want: %#v", todo.ID, id)
	}
	if todo.Body != body {
		t.Errorf("\nUnexpected body got: %#v want: %#v", todo.Body, body)
	}
}

func TestGetTodo(t *testing.T) {
	tApp := App{Database: &db.TestDatabase{}}
	id := 1
	todo, err := tApp.GetTodo(uint(id))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if todo.ID != uint(id) {
		t.Errorf("Unexpected id got: %#v want: %#v", todo.ID, id)
	}
}

func TestGetTodos(t *testing.T) {
	tTodos := &[]model.Todo{
		{Body: "Test1"},
		{Body: "Test2"},
		{Body: "Test3"},
	}
	tDb := &db.TestDatabase{Todos: tTodos}
	tApp := App{Database: tDb}

	todos, err := tApp.GetTodos()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(tTodos, todos) {
		t.Errorf("\nUnexpected result got:\n %+v \nwant:\n %+v", todos, tTodos)
	}
}

func TestGetActiveTodos(t *testing.T) {
	tTodos := &[]model.Todo{
		{Body: "Test1", Done: false},
		{Body: "Test2", Done: false},
		{Body: "Test3", Done: false},
	}
	tDb := &db.TestDatabase{Todos: tTodos}
	tApp := App{Database: tDb}

	todos, err := tApp.GetActiveTodos()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	for index, todo := range *todos {
		if todo.Done != false {
			t.Errorf("todo[%v] should be active, got: todo.Done = %v, want: todo.Done = %v", index, todo.Done, false)
		}
	}
}

func TestGetDoneTodos(t *testing.T) {
	tTodos := &[]model.Todo{
		{Body: "Test1", Done: true},
		{Body: "Test2", Done: true},
		{Body: "Test3", Done: true},
	}
	tDb := &db.TestDatabase{Todos: tTodos}
	tApp := App{Database: tDb}

	todos, err := tApp.GetDoneTodos()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	for index, todo := range *todos {
		if todo.Done != true {
			t.Errorf("todo[%v] should be done, got: todo.Done = %v, want: todo.Done = %v", index, todo.Done, true)
		}
	}
}

func TestDeleteTodo(t *testing.T) {
	tApp := App{Database: &db.TestDatabase{}}
	if err := tApp.DeleteTodo(1); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestDoneTodo(t *testing.T) {
	tApp := App{Database: &db.TestDatabase{}}
	id := 1
	todo, err := tApp.DoneTodo(uint(id))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if todo.ID != uint(id) {
		t.Errorf("Unexpected todo id got: %#v want: %#v", todo.ID, id)
	}
	if todo.Done != true {
		t.Errorf("\ntodo should be done, got: todo.Done = %v, want: todo.Done = %v", todo.Done, true)
	}
}

func TestUndoneTodo(t *testing.T) {
	tApp := App{Database: &db.TestDatabase{}}
	id := 1
	todo, err := tApp.UndoneTodo(uint(id))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if todo.ID != uint(id) {
		t.Errorf("Unexpected todo id got: %#v want: %#v", todo.ID, id)
	}
	if todo.Done != false {
		t.Errorf("\ntodo should be undone, got: todo.Done = %v, want: todo.Done = %v", todo.Done, false)
	}
}
