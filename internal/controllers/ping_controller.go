package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
