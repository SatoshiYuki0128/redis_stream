package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var GinRouter *gin.Engine

func InitGin() {
	GinRouter = gin.Default()

	GinRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin web service!",
		})
	})
}
