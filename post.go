package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func publishPost(c *gin.Context) {
	publishPostHelper(c, nil)
}

func publishChildPost(c *gin.Context) {
	idParam := c.Param("id")

	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	id := uint(idInt)

	publishPostHelper(c, &id)
}

func publishPostHelper(c *gin.Context, id *uint) {
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
	if id != nil {
		post.ParentId = id
	}
	app.dal.createPost(&post)

	categories := app.dal.db.Model(&post).Association("Categories")

	for _, categoryName := range newPost.Categories {
		category := app.dal.getCategoryByTitle(categoryName)
		categories.Append(category)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
