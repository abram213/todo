package app

import (
	"fmt"
	"todo/db"
	"todo/model"
)

type App struct {
	Config   *Config
	Database db.DataLayer
}

func New(dataMode string) (app *App, err error) {
	app = &App{}
	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	if dataMode == "" {
		dataMode = "sql"
	}
	if err := app.initDatabase(dataMode); err != nil {
		return nil, err
	}
	return app, nil
}

func (app *App) initDatabase(dataMode string) error {
	dbConfig, err := db.InitConfig(dataMode)
	if err != nil {
		return err
	}
	switch dataMode {
	case "sql":
		if err := app.initSQLDatabase(dbConfig); err != nil {
			return err
		}
	case "file":
		if err := app.initFileDatabase(dbConfig); err != nil {
			return err
		}
	case "mongo":
		if err := app.initMongoDatabase(dbConfig); err != nil {
			return err
		}
	default:
		return fmt.Errorf("not supported db mode: %v", dataMode)
	}
	return nil
}

func (app *App) initSQLDatabase(config *db.Config) (err error) {
	app.Database, err = db.NewSQLDatabase(config)
	if err != nil {
		return
	}
	app.Database.Migrate(
		&model.Todo{})
	return
}

func (app *App) initFileDatabase(config *db.Config) (err error) {
	app.Database, err = db.NewFileDatabase(config)
	if err != nil {
		return
	}
	return
}

func (app *App) initMongoDatabase(config *db.Config) (err error) {
	app.Database, err = db.NewMongoDatabase(config)
	if err != nil {
		return
	}
	return
}
