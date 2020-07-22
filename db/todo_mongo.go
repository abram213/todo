package db

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
	"todo/errs"
	"todo/model"
)

func (db *MongoDatabase) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	todo.UpdatedAt = time.Now()

	collection := db.Database("todo").Collection("todos")
	if _, err := collection.InsertOne(context.TODO(), todo); err != nil {
		return nil, errors.Wrap(err, "unable to create todo")
	}
	return todo, nil
}

func (db *MongoDatabase) UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	filter := bson.D{{"model.id", todo.ID}}
	collection := db.Database("todo").Collection("todos")
	update := bson.D{
		{"$set", bson.D{
			{"body", todo.Body},
			{"done", todo.Done},
			{"done_time", todo.DoneTime},
			{"model.updated_at", time.Now()},
		}},
	}
	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errs.CustomError{
				Message: fmt.Sprintf("cant find todo with id: %v", todo.ID),
				Code:    http.StatusBadRequest,
			}
		}
		return nil, errors.Wrap(err, "unable to update todo")
	}
	return todo, nil
}

func (db *MongoDatabase) GetTodo(id uint) (*model.Todo, error) {
	filter := bson.D{{"model.id", id}}
	collection := db.Database("todo").Collection("todos")
	var todo model.Todo
	if err := collection.FindOne(context.TODO(), filter).Decode(&todo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errs.CustomError{
				Message: fmt.Sprintf("cant find todo with id: %v", id),
				Code:    http.StatusBadRequest,
			}
		}
		return nil, errors.Wrap(err, "unable to get todo")
	}
	return &todo, nil
}

func (db *MongoDatabase) GetTodos() (*[]model.Todo, error) {
	collection := db.Database("todo").Collection("todos")
	opt := options.Find()

	var todos []model.Todo

	cur, err := collection.Find(context.TODO(), bson.M{}, opt)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get todos")
	}
	if err := cur.All(context.TODO(), &todos); err != nil {
		return nil, errors.Wrap(err, "unable to get todos")
	}
	return &todos, nil
}

func (db *MongoDatabase) GetActiveTodos() (*[]model.Todo, error) {
	collection := db.Database("todo").Collection("todos")

	opt := options.Find()
	filter := bson.D{{"done", false}}

	var todos []model.Todo
	cur, err := collection.Find(context.TODO(), filter, opt)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get todos")
	}
	if err := cur.All(context.TODO(), &todos); err != nil {
		return nil, errors.Wrap(err, "unable to get todos")
	}
	return &todos, nil
}

func (db *MongoDatabase) GetDoneTodos() (*[]model.Todo, error) {
	collection := db.Database("todo").Collection("todos")
	opt := options.Find()
	filter := bson.D{{"done", true}}

	var todos []model.Todo
	cur, err := collection.Find(context.TODO(), filter, opt)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get todos")
	}
	if err := cur.All(context.TODO(), &todos); err != nil {
		return nil, errors.Wrap(err, "unable to get todos")
	}
	return &todos, nil
}

func (db *MongoDatabase) DeleteTodo(id uint) error {
	filter := bson.D{{"model.id", id}}
	collection := db.Database("todo").Collection("todos")
	if _, err := collection.DeleteOne(context.TODO(), filter); err != nil {
		return err
	}
	return nil
}
