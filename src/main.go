package main

import (
	"net/http"

	engine "github.com/edista/go-microservice/engine"
	"github.com/edista/go-microservice/engine/src/database"
	v1 "github.com/edista/go-microservice/src/api/v1"
	"github.com/edista/go-microservice/src/config"
	"github.com/edista/go-microservice/src/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	app := echo.New()

	config, err := config.NewConfig()
	if err != nil {

	}

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Health Check
	app.GET("/c/health", hello)

	db, err := database.NewDatabase(config.Database.Uri)
	if err != nil {
		return
	}
	engine := engine.NewEngine(app, db)
	service := service.NewService(engine, config)
	v1.NewApi(app, service)

	// Start server
	app.Logger.Fatal(app.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
