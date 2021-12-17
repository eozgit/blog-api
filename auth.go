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

	authorized := createProtectedRouterGroup()

	createProtectedEndpoints(authorized)

	c.JSON(http.StatusOK, gin.H{"status": "User created"})
}

func getAccounts() map[string]string {
	users := app.dal.listUsers()

	accounts := make(map[string]string)
	for i := 0; i < len(users); i++ {
		user := users[i]
		accounts[user.Username] = user.Password
	}

	return accounts
}

func createProtectedRouterGroup() *gin.RouterGroup {
	accounts := getAccounts()

	return app.r.Group("/", gin.BasicAuth(accounts))
}

func checkAuthorisation(c *gin.Context) (string, bool) {
	username, _, ok := c.Request.BasicAuth()

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised"})
	}

	return username, ok
}
