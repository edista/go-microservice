package engine

import (
	"github.com/edista/go-microservice/engine/src/database"
	"github.com/labstack/echo/v4"
)

type Engine struct {
	App      *echo.Echo
	Database *database.Database
}

func NewEngine(app *echo.Echo, database *database.Database) *Engine {
	return &Engine{
		App:      app,
		Database: database,
	}
}
