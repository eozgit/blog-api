package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func publishPost(c *gin.Context) {
	username, ok := checkAuthorisation(c)
	if !ok {
		return
	}

	publishPostHelper(c, nil, username)
}

func publishChildPost(c *gin.Context) {
	username, ok := checkAuthorisation(c)
	if !ok {
		return
	}

	id, hasErr := getIdParam(c)
	if hasErr {
		return
	}

	publishPostHelper(c, id, username)
}

func publishPostHelper(c *gin.Context, id *uint, username string) {
	var newPost NewPost
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, gin.H{"status": "Post created"})
}

func updatePost(c *gin.Context) {
	username, ok := checkAuthorisation(c)
	if !ok {
		return
	}

	id, hasErr := getIdParam(c)
	if hasErr {
		return
	}

	user := app.dal.getUserByName(username)

	post := app.dal.getPostByID(id)

	if user.ID != post.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised"})
		return
	}

	var updatedPost UpdatedPost
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Title = updatedPost.Title
	post.Content = updatedPost.Content
	app.dal.db.Save(&post)

	c.JSON(http.StatusOK, gin.H{"status": "Post updated"})
}
