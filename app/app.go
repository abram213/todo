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

func New() (app *App, err error) {
	app = &App{}
	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	if err := app.initDatabase(); err != nil {
		return nil, err
	}
	return app, nil
}

func (app *App) initDatabase() error {
	dbConfig, err := db.InitConfig()
	if err != nil {
		return err
	}
	switch app.Config.DataMode {
	case "sql":
		if err := app.initSQLDatabase(dbConfig); err != nil {
			return err
		}
	case "file":
		if err := app.initFileDatabase(dbConfig); err != nil {
			return err
		}
	default:
		return fmt.Errorf("not supported db mode: %v", app.Config.DataMode)
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
