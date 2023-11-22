package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_const "redis_stream/const"
	"redis_stream/handler"
)

var GinRouter *gin.Engine

func InitGin() error {
	GinRouter = gin.Default()

	err := GinRouter.SetTrustedProxies([]string{_const.LocalHost})
	if err != nil {
		return err
	}

	GinRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin web service!",
		})
	})

	GinRouter.PUT("/stream", handler.CreateStream)
	GinRouter.DELETE("/stream/:id", handler.DeleteStream)
	GinRouter.GET("/stream", handler.GetStreamList)

	return nil
}
