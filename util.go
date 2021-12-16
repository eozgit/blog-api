package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getIdParam(c *gin.Context) (*uint, bool) {
	idParam := c.Param("id")

	idInt, err := strconv.Atoi(idParam)
	hasErr := err != nil
	var id *uint = nil
	if hasErr {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint("invalid id %w %w", idParam, err.Error())})
	} else {
		idHelper := uint(idInt)
		id = &idHelper
	}

	return id, hasErr
}
