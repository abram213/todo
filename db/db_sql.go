package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

type SQLDatabase struct {
	*gorm.DB
}

func NewSQLDatabase(config *Config) (*SQLDatabase, error) {
	db, err := gorm.Open("sqlite3", config.DatabaseURI)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &SQLDatabase{db}, nil
}

func (db *SQLDatabase) CloseDB() error {
	return db.Close()
}

func (db *SQLDatabase) Migrate(values ...interface{}) {
	db.AutoMigrate(values...)
}

func (db *SQLDatabase) DropTables(values ...interface{}) {
	db.DropTableIfExists(values...)
}
