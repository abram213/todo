package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"todo/app"
	"todo/model"
)

func (a *API) todoRouter() http.Handler {
	r := chi.NewRouter()
	r.Method("GET", "/", a.handler(a.GetTodos))
	r.Method("GET", "/active", a.handler(a.GetActiveTodos))
	r.Method("GET", "/done", a.handler(a.GetDoneTodos))
	r.Method("POST", "/", a.handler(a.CreateTodo))
	r.Route("/{id:^[0-9]*$}", func(r chi.Router) {
		r.Method("GET", "/", a.handler(a.GetTodo))
		r.Method("PUT", "/", a.handler(a.UpdateTodo))
		r.Method("DELETE", "/", a.handler(a.DeleteTodo))
		r.Method("GET", "/done", a.handler(a.DoneTodo))
		r.Method("GET", "/undone", a.handler(a.UndoneTodo))
	})
	return r
}

type todoInput struct {
	Body string `json:"body"`
}

func (a *API) GetTodos(w http.ResponseWriter, r *http.Request) error {
	todos, err := a.App.GetTodos()
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(todos)
	return nil
}

func (a *API) GetActiveTodos(w http.ResponseWriter, r *http.Request) error {
	todos, err := a.App.GetActiveTodos()
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(todos)
	return nil
}

func (a *API) GetDoneTodos(w http.ResponseWriter, r *http.Request) error {
	todos, err := a.App.GetDoneTodos()
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(todos)
	return nil
}

func (a *API) CreateTodo(w http.ResponseWriter, r *http.Request) error {
	var input todoInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &input); err != nil {
		return &app.CustomError{
			Message: errors.Wrap(err, "parse json error").Error(),
			Code:    http.StatusBadRequest,
		}
	}

	todo := &model.Todo{
		Body: input.Body,
	}
	if _, err := a.App.CreateTodo(todo); err != nil {
		return err
	}
	return nil
}

func (a *API) GetTodo(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	todo, err := a.App.GetTodo(uint(uid))
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(todo)
	return nil
}

func (a *API) UpdateTodo(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	var input todoInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &input); err != nil {
		return &app.CustomError{
			Message: errors.Wrap(err, "parse json error").Error(),
			Code:    http.StatusBadRequest,
		}
	}

	todo, err := a.App.UpdateTodo(uint(uid), input.Body)
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(todo)
	return nil
}

func (a *API) DeleteTodo(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	if err := a.App.DeleteTodo(uint(uid)); err != nil {
		return err
	}
	return nil
}

func (a *API) DoneTodo(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	todo, err := a.App.DoneTodo(uint(uid))
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(todo)
	return nil
}

func (a *API) UndoneTodo(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	todo, err := a.App.UndoneTodo(uint(uid))
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(todo)
	return nil
}
