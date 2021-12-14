package main

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	dal *DataAccessLayer
	r   *gin.Engine
}

var app *App

func main() {
	app = setupApp()

	app.r.Run()
}

func createProtectedEndpoints(authorized *gin.RouterGroup) {
	authorized.POST("/post", publishPost)
	authorized.POST("/post/:id", publishChildPost)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", registerUser)
	return r
}

func setupApp() *App {
	app = &App{&DataAccessLayer{}, setupRouter()}
	app.dal.init()

	app.r = setupRouter()
	return app
}
