package app

type ServerInterface interface {
	GetServer(*App)
}

type Server struct {
}

func (s *Server) GetServer(app *App) {
	//app.Echo.GET("/cars", app.HandleGetAllAutomobiles)
	//app.Echo.GET("/carById", app.HandleGetAutomobileById)
	//app.Echo.GET("/InspectionById", app.HandleGetInspectionsById)
}
