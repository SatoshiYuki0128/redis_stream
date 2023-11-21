package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var GinRouter *gin.Engine

func InitGin() error {
	GinRouter = gin.Default()

	// 設置信任的代理
	err := GinRouter.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return err
	}

	GinRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin web service!",
		})
	})

	return nil
}
