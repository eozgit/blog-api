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
	dal.createUser(&user)

	authorized := createProtectedRouterGroup()

	createProtectedEndpoints(authorized)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func getAccounts() map[string]string {
	users := dal.listUsers()

	accounts := make(map[string]string)
	for i := 0; i < len(users); i++ {
		user := users[i]
		accounts[user.Username] = user.Password
	}

	return accounts
}

func createProtectedRouterGroup() *gin.RouterGroup {
	accounts := getAccounts()

	return r.Group("/", gin.BasicAuth(accounts))
}
