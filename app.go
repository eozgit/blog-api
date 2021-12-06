package main

import "github.com/gin-gonic/gin"

var dal DataAccessLayer

func main() {
	dal = DataAccessLayer{}
	dal.init()

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {

		c.String(200, "ok")
	})
	r.Run()
}
