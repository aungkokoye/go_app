package services

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/aungkokoye/go_app/databases"
	"github.com/aungkokoye/go_app/repositiories"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppInterface interface {
	GetDatabase(db string) (*sql.DB, error)
}

type App struct {
	Config   Config
	database map[string]*sql.DB
}

func NewApp(config Config) *App {

	setLoggingLevel(config.LogLevel)

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

func (a *App) RunGin() {
	router := gin.Default()

	router.GET("/albums", repositiories.GetAlbums)
	router.Run("localhost:8080")
	fmt.Println("Running go_api_gin!")
}

func setLoggingLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

}
