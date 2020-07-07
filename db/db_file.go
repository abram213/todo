package db

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"todo/model"
)

type fData struct {
	todos  map[uint]*model.Todo
	lastID uint
}

type FileDatabase struct {
	path string
	file *os.File
	data fData
}

func NewFileDatabase(config *Config) (*FileDatabase, error) {
	file, err := os.OpenFile(config.PathToFile, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	todos, lastID, err := parseTodos(file)
	if err != nil {
		return nil, err
	}
	return &FileDatabase{config.PathToFile, file, fData{todos: todos, lastID: lastID}}, nil
}

func (fdb *FileDatabase) CloseDB() error {
	if err := fdb.saveTodos(); err != nil {
		return err
	}
	return nil
}
func (fdb *FileDatabase) Migrate(values ...interface{})    {}
func (fdb *FileDatabase) DropTables(values ...interface{}) {}

func parseTodos(file *os.File) (map[uint]*model.Todo, uint, error) {
	todos := map[uint]*model.Todo{}
	reader := bufio.NewReader(file)

	var line string
	var rErr error
	var id uint
	for {
		line, rErr = reader.ReadString('\n')
		todo, err := parseTodo(strings.TrimSpace(line))
		if err == nil {
			todos[todo.ID] = &todo
			if todo.ID > id {
				id = todo.ID
			}
		}

		if rErr != nil {
			break
		}
	}

	if rErr != io.EOF {
		return nil, 0, rErr
	}
	return todos, id, nil
}

func parseTodo(line string) (model.Todo, error) {
	var todo model.Todo
	parts := strings.Split(line, "|")
	if len(parts) != 5 {
		return todo, fmt.Errorf("bad string format: %v", line)
	}
	id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return todo, err
	}
	updateTime, err := time.Parse("02-01-2006/15:04:05", strings.TrimSpace(parts[1]))
	if err != nil {
		return todo, err
	}
	done := strings.TrimSpace(parts[3]) == "true"
	doneTime, err := time.Parse("02-01-2006/15:04:05", strings.TrimSpace(parts[4]))
	if err != nil {
		return todo, err
	}

	todo.Model.ID = uint(id)
	todo.Model.UpdatedAt = updateTime
	todo.Body = strings.TrimSpace(parts[2])
	todo.Done = done
	todo.DoneTime = doneTime
	return todo, nil
}
