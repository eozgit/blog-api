package main

import (
	"github.com/gin-gonic/gin"
)

var dal DataAccessLayer

func main() {
	dal = DataAccessLayer{}
	dal.init()

	accounts := getAccounts()

	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(accounts))

	authorized.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.Run()
}

func getAccounts() map[string]string {
	users := dal.listUsers()

	accounts := make(map[string]string)
	for i := 0; i < len(users); i++ {
		user := users[i]
		accounts[user.UserName] = user.PasswordHash
	}

	return accounts
}
