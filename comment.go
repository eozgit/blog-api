package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func publishComment(c *gin.Context) {
	publishCommentHelper(c, false)
}

func publishChildComment(c *gin.Context) {
	publishCommentHelper(c, true)
}

func publishCommentHelper(c *gin.Context, isChildComment bool) {
	username, ok := checkAuthorisation(c)
	if !ok {
		return
	}

	id, hasErr := getIdParam(c)
	if hasErr {
		return
	}

	var newComment NewComment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := app.dal.getUserByName(username)

	comment := Comment{}
	comment.Title = newComment.Title
	comment.Content = newComment.Content
	comment.User = *user
	if isChildComment {
		comment.ParentId = id
	} else {
		comment.PostID = id
	}
	app.dal.createComment(&comment)

	c.IndentedJSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
