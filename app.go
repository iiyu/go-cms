package main

import (
	"github.com/cihub/seelog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	engine          *gin.Engine
	db              *gorm.DB
	responseHandler ResponseHandler
	requestHandler  RequestHandler
}

func InitApp() *App {
	seelog.ReplaceLogger(Logger)
	defer seelog.Flush()
	db, err := gorm.Open("mysql", Conn)
	if err != nil {
		seelog.Critical("Could not connect database", err)
	}
	db.LogMode(true)
	db.SingularTable(true)

	// Migrate the schema
	db.AutoMigrate(
		&User{},
	)

	responseHandler := NewResponseHandler()
	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(responseHandler.NoRoute)

	app := &App{
		r,
		db,
		responseHandler,
		NewRequestHandler(),
	}

	InitHandlers(app)

	return app
}

func (app *App) Run() {
	app.Engine().Run(":8077")
}

func (app *App) Engine() *gin.Engine {
	return app.engine
}

func (app *App) Db() *gorm.DB {
	return app.db
}

func (app *App) ResponseHandler() ResponseHandler {
	return app.responseHandler
}

func (app *App) RequestHandler() RequestHandler {
	return app.requestHandler
}
