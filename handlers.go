package main

func InitHandlers(app *App) {
	InitAuthHandler(app)
	InitClubsHandler(app)
	InitTablesHandler(app)
}
