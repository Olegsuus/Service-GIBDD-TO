package handlers

import "github.com/labstack/echo/v4"

type Handlers interface {
	RegisterRoutes(e *echo.Echo)
}

//func (a *app.App.DB) HandleGetAutomobileById(c echo.Context) error {
//	strId := c.QueryParam("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
//	}
//
//	automobile, err := a.DB.GetAutomobileById(id)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, automobile)
//}
//
//func (a *App) HandleGetInspectionsById(c echo.Context) error {
//	strId := c.QueryParam("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
//	}
//
//	inspection, err := a.DB.GetInspectionsById(id)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, inspection)
//}
//
//func (a *App) HandleGetAllAutomobiles(c echo.Context) error {
//	automobile, err := a.DB.GetAllAutomobiles()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"Error": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, automobile)
//}
//
//func (a *App) HandleAddNewAutomobile(c echo.Context) error {
//	var automobile *models.Automobile
//	if err := c.Bind(&automobile); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
//	}
//
//	err := a.DB.AddNewAutomobile(automobile)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
//	}
//
//	return c.JSON(http.StatusCreated, automobile)
//}
