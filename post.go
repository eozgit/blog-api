package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func publishPost(c *gin.Context) {
	var newPost NewPost
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _, ok := c.Request.BasicAuth()

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unauthorised"})
		return
	}

	user := app.dal.getUserByName(username)

	post := Post{}
	post.Title = newPost.Title
	post.Content = newPost.Content
	post.User = *user
	app.dal.createPost(&post)

	categories := app.dal.db.Model(&post).Association("Categories")

	for _, categoryName := range newPost.Categories {
		category := app.dal.getCategoryByTitle(categoryName)
		categories.Append(category)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
