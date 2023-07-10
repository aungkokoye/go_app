package services

import (
	"database/sql"

	"github.com/aungkokoye/go_app/databases"
)

type AppInterface interface {
	GetDatabase(db string) (*sql.DB, error)
}

type App struct {
	Config   Config
	database map[string]*sql.DB
}

func NewApp(config Config) *App {

	return &App{
		Config:   config,
		database: make(map[string]*sql.DB),
	}

}

// GetDatabase returns a database connection for the given database name
func (a *App) GetDatabase(db string) (*sql.DB, error) {
	if _, ok := a.database[db]; !ok {
		var err error
		a.database[db], err = databases.NewMysqlDB(db, a.Config.MysqlDB)
		if err != nil {
			return nil, err
		}
	}

	return a.database[db], nil
}
