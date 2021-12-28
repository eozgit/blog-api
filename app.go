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

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", registerUser)
	r.POST("/post", makeSecure(publishPost))
	r.GET("/post", listPostsByCategory)
	r.POST("/post/:id", makeSecure(publishChildPost))
	r.PATCH("/post/:id", makeSecure(updatePost))
	r.POST("/post/:id/comment", makeSecure(publishComment))
	r.GET("/post/:id/comment", listCommentsByPost)
	r.POST("/comment/:id", makeSecure(publishChildComment))
	r.DELETE("/comment/:id", makeSecure(deleteComment))
	return r
}

func setupApp() *App {
	app = &App{&DataAccessLayer{}, setupRouter()}
	app.dal.init()

	app.r = setupRouter()
	return app
}
