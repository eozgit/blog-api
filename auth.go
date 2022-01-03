package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUser(c *gin.Context) {
	var credentials Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{}
	user.Username = credentials.Username
	user.Password = credentials.Password
	app.dal.createUser(&user)

	c.JSON(http.StatusOK, gin.H{"status": "User created"})
}

func makeSecure(handler func(*gin.Context)) func(*gin.Context) {
	return func(c *gin.Context) {
		username, password, _ := c.Request.BasicAuth()

		user := app.dal.getUserByName(username)

		if user.Password != password {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised"})
		} else {
			handler(c)
		}
	}
}

func getUsername(c *gin.Context) string {
	username, _, _ := c.Request.BasicAuth()

	return username
}
