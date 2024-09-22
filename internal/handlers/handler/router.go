package handler

import "github.com/labstack/echo/v4"

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.Static("/", "frontend")

	e.GET("/cars", h.AutomobileHandler.GetAll)
	e.GET("/car/:id", h.AutomobileHandler.Get)
	e.POST("/car", h.AutomobileHandler.Add)
	e.PATCH("/car/:id", h.AutomobileHandler.Update)
	e.DELETE("/car/:id", h.AutomobileHandler.Delete)
	e.GET("/report", h.AutomobileHandler.GetReport)

	e.GET("/inspections", h.InspectionHandler.GetAll)
	e.GET("/inspection/:id", h.InspectionHandler.Get)
	e.POST("/inspection", h.InspectionHandler.Add)
	e.PATCH("/inspection/:id", h.InspectionHandler.Update)
	e.DELETE("/inspection/:id", h.InspectionHandler.Delete)
}
