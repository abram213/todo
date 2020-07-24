package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"todo/db"
	"todo/model"
)

func TestCreateTodoHandler(t *testing.T) {
	router, err := NewTestRouter(&db.TestDatabase{})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	ts := httptest.NewServer(router)
	defer ts.Close()

	tTodo := model.Todo{
		Body: "Test",
	}
	path := "/api/todo"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(tTodo)
	req, body, err := testRequest(t, ts, "POST", path, buf)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	if req.StatusCode != http.StatusOK {
		t.Errorf("Invalid status code: %d want: %d", req.StatusCode, http.StatusOK)
	}

	var todo model.Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(tTodo, todo) {
		t.Errorf("\nUnexpected result got:\n %+v \nwant:\n %+v", todo, tTodo)
	}
}

func TestUpdateTodoHandler(t *testing.T) {
	router, err := NewTestRouter(&db.TestDatabase{})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	ts := httptest.NewServer(router)
	defer ts.Close()

	tTodo := model.Todo{
		Body: "TestTodo",
	}
	id := 1
	path := "/api/todo/" + strconv.Itoa(id)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(tTodo)
	req, body, err := testRequest(t, ts, "PUT", path, buf)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	if req.StatusCode != http.StatusOK {
		t.Errorf("Invalid status code: %d want: %d", req.StatusCode, http.StatusOK)
	}

	var todo model.Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if todo.ID != uint(id) {
		t.Errorf("\nUnexpected id got: %#v want: %#v", todo.ID, id)
	}
	if todo.Body != tTodo.Body {
		t.Errorf("\nUnexpected body got: %#v want: %#v", todo.Body, tTodo.Body)
	}
}

func TestGetTodoHandler(t *testing.T) {
	router, err := NewTestRouter(&db.TestDatabase{})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	ts := httptest.NewServer(router)
	defer ts.Close()

	id := 1
	path := "/api/todo/" + strconv.Itoa(id)

	_, body, err := testRequest(t, ts, "GET", path, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	var todo model.Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if todo.ID != uint(id) {
		t.Errorf("Invalid todo id %d expected %d", todo.ID, id)
	}
}

func TestGetTodosHandler(t *testing.T) {
	tTodos := []model.Todo{
		{Body: "Test1"},
		{Body: "Test2"},
		{Body: "Test3"},
	}
	tDb := &db.TestDatabase{Todos: &tTodos}

	router, err := NewTestRouter(tDb)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	ts := httptest.NewServer(router)
	defer ts.Close()

	path := "/api/todo"

	_, body, err := testRequest(t, ts, "GET", path, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	var todos []model.Todo
	if err := json.Unmarshal(body, &todos); err != nil {
		t.Errorf("Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(tTodos, todos) {
		t.Errorf("\nUnexpected result got:\n %+v \nwant:\n %+v", todos, tTodos)
	}
}

func TestGetActiveTodosHandler(t *testing.T) {
	tTodos := []model.Todo{
		{Body: "Test1", Done: false},
		{Body: "Test2", Done: false},
		{Body: "Test3", Done: false},
	}
	tDb := &db.TestDatabase{Todos: &tTodos}

	router, err := NewTestRouter(tDb)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	ts := httptest.NewServer(router)
	defer ts.Close()

	path := "/api/todo/active"

	_, body, err := testRequest(t, ts, "GET", path, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	var todos []model.Todo
	if err := json.Unmarshal(body, &todos); err != nil {
		t.Errorf("Unmarshal error: %v", err)
	}

	for index, todo := range todos {
		if todo.Done != false {
			t.Errorf("todo[%v] should be active, got: todo.Done = %v, want: todo.Done = %v", index, todo.Done, false)
		}
	}
}

func TestGetDoneTodosHandler(t *testing.T) {
	tTodos := []model.Todo{
		{Body: "Test1", Done: true},
		{Body: "Test2", Done: true},
		{Body: "Test3", Done: true},
	}
	tDb := &db.TestDatabase{Todos: &tTodos}

	router, err := NewTestRouter(tDb)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	ts := httptest.NewServer(router)
	defer ts.Close()

	path := "/api/todo/done"

	_, body, err := testRequest(t, ts, "GET", path, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	var todos []model.Todo
	if err := json.Unmarshal(body, &todos); err != nil {
		t.Errorf("Unmarshal error: %v", err)
	}

	for index, todo := range todos {
		if todo.Done != true {
			t.Errorf("todo[%v] should be active, got: todo.Done = %v, want: todo.Done = %v", index, todo.Done, true)
		}
	}
}

func TestDeleteTodoHandler(t *testing.T) {
	router, err := NewTestRouter(&db.TestDatabase{})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	ts := httptest.NewServer(router)
	defer ts.Close()

	id := 3
	path := "/api/todo/" + strconv.Itoa(id)

	req, _, err := testRequest(t, ts, "DELETE", path, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	if req.StatusCode != http.StatusOK {
		t.Errorf("Invalid status code: %d want: %d", req.StatusCode, http.StatusOK)
	}
}

func TestDoneTodoHandler(t *testing.T) {
	router, err := NewTestRouter(&db.TestDatabase{})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	ts := httptest.NewServer(router)
	defer ts.Close()

	id := 1
	path := "/api/todo/" + strconv.Itoa(id) + "/done"

	_, body, err := testRequest(t, ts, "GET", path, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	var todo model.Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if todo.ID != uint(id) {
		t.Errorf("Unexpected todo id got: %#v want: %#v", todo.ID, id)
	}
	if todo.Done != true {
		t.Errorf("\ntodo should be done, got: todo.Done = %v, want: todo.Done = %v", todo.Done, true)
	}
}

func TestUndoneTodoHandler(t *testing.T) {
	router, err := NewTestRouter(&db.TestDatabase{})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	ts := httptest.NewServer(router)
	defer ts.Close()

	id := 1
	path := "/api/todo/" + strconv.Itoa(id) + "/undone"

	_, body, err := testRequest(t, ts, "GET", path, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	var todo model.Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if todo.ID != uint(id) {
		t.Errorf("Unexpected todo id got: %#v want: %#v", todo.ID, id)
	}
	if todo.Done != false {
		t.Errorf("\ntodo should be done, got: todo.Done = %v, want: todo.Done = %v", todo.Done, false)
	}
}
