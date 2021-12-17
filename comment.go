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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, gin.H{"status": "Comment created"})
}

func deleteComment(c *gin.Context) {
	username, ok := checkAuthorisation(c)
	if !ok {
		return
	}

	id, hasErr := getIdParam(c)
	if hasErr {
		return
	}

	user := app.dal.getUserByName(username)

	comment := app.dal.getCommentById(*id)

	if user.ID != comment.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised"})
		return
	}

	app.dal.deleteCommentById(*id)

	c.JSON(http.StatusOK, gin.H{"status": "Comment deleted"})
}
