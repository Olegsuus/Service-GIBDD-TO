package app

import (
	"fmt"
	"github.com/Olegsuus/TZ-WEB-App/internal/config"
	handler_automobile "github.com/Olegsuus/TZ-WEB-App/internal/handlers/automobile"
	"github.com/Olegsuus/TZ-WEB-App/internal/handlers/handler"
	handler_inspection "github.com/Olegsuus/TZ-WEB-App/internal/handlers/inspection"
	service_automobile "github.com/Olegsuus/TZ-WEB-App/internal/service/automobile"
	service_inspection "github.com/Olegsuus/TZ-WEB-App/internal/service/inspection"
	"github.com/Olegsuus/TZ-WEB-App/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
)

type App struct {
	Config   *config.Config
	Handlers *handler.Handler
	Echo     *echo.Echo
	Logger   *slog.Logger
}

func NewApp(cfg *config.Config, storage *storage.Storage, logger *slog.Logger) *App {
	return &App{
		Config: cfg,
		Logger: logger,
	}
}

func (a *App) InitializeHandlers(store *storage.Storage) {
	autoService := service_automobile.NewAutomobileService(store.AutomobileStorage, a.Logger)
	inspService := service_inspection.NewInspectionService(store.InspectionStorage, a.Logger)

	autoHandler := handler_automobile.NewAutomobileHandler(autoService)
	inspHandler := handler_inspection.NewInspectionHandler(inspService)

	a.Handlers = handler.NewHandler(autoHandler, inspHandler)
}

func (a *App) Start(store *storage.Storage) error {
	a.Echo = echo.New()

	a.Echo.Use(middleware.Logger())
	a.Echo.Use(middleware.Recover())

	a.InitializeHandlers(store)

	a.Handlers.RegisterRoutes(a.Echo)

	a.registerSwaggerRoutes()

	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	a.Logger.Info("Starting server", "address", addr)

	return a.Echo.Start(addr)
}

func (a *App) Stop() error {
	return nil
}

func (a *App) registerSwaggerRoutes() {
	a.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
