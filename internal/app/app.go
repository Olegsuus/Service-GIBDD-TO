package app

import (
	"fmt"
	"log"

	"Web-App/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	Config          *config.Config
	DB              Storage
	ServerInterface ServerInterface
	Echo            *echo.Echo
}

type Storage interface {
	GetStorage(cfg *config.Config)
	Stop() error
	//GetAutomobileById(id int) (*models.Automobile, error)
	//GetInspectionsById(id int) (*models.Inspection, error)
	//GetAllAutomobiles() ([]models.Automobile, error)
}

func (a *App) Start() error {
	a.Echo = echo.New()
	a.ServerInterface.GetServer(a)
	a.Echo.Use(middleware.Logger())
	a.Echo.Use(middleware.Recover())

	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	log.Printf("Starting server on %s", addr)
	return a.Echo.Start(addr)
}

func (a *App) Stop() {
	//
}
