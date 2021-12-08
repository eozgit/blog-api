package main

import (
	"github.com/gin-gonic/gin"
)

var dal DataAccessLayer
var r *gin.Engine

func main() {
	dal = DataAccessLayer{}
	dal.init()

	r = gin.Default()

	r.POST("/register", registerUser)

	r.Run()
}

func createProtectedEndpoints(authorized *gin.RouterGroup) {
	authorized.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})
}
