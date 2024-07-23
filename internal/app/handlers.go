package app

//func (a *App) HandleGetAutomobileById(c echo.Context) error {
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
