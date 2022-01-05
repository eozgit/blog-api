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
	id, hasErr := getIdParam(c)
	if hasErr {
		return
	}

	var newComment NewComment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := getUsername(c)
	user := app.dal.getUserByName(username)

	comment := Comment{}
	comment.Title = newComment.Title
	comment.Content = newComment.Content
	comment.User = *user
	if isChildComment {
		parentComment := app.dal.getCommentById(*id)
		comment.PostID = parentComment.PostID
		comment.ParentId = id
	} else {
		comment.PostID = id
	}
	app.dal.createComment(&comment)

	c.JSON(http.StatusOK, gin.H{"status": "Comment created"})
}

func deleteComment(c *gin.Context) {
	id, hasErr := getIdParam(c)
	if hasErr {
		return
	}

	username := getUsername(c)
	user := app.dal.getUserByName(username)

	comment := app.dal.getCommentById(*id)

	if user.ID != comment.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised"})
		return
	}

	app.dal.deleteCommentById(*id)

	c.JSON(http.StatusOK, gin.H{"status": "Comment deleted"})
}

func listCommentsByPost(c *gin.Context) {
	id, hasErr := getIdParam(c)
	if hasErr {
		return
	}

	post := app.dal.getPostByID(id)

	comments := app.dal.listCommentsByPost(post)

	c.JSON(http.StatusOK, comments)
}
