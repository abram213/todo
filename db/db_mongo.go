package db

import (
	"context"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	*mongo.Client
}

func NewMongoDatabase(config *Config) (*MongoDatabase, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DatabaseURI))
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &MongoDatabase{client}, nil
}

func (db *MongoDatabase) CloseDB() error {
	return db.Disconnect(context.TODO())
}

func (db *MongoDatabase) Migrate(values ...interface{}) {}

func (db *MongoDatabase) DropTables(values ...interface{}) {}
