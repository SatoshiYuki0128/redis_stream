package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_const "redis_stream/const"
	"redis_stream/handler"
)

var GinRouter *gin.Engine

func InitGin() error {
	//gin.SetMode(gin.ReleaseMode)
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

	v1 := GinRouter.Group("/api/v1")
	{
		v1.PUT("/stream", handler.CreateStream)
		v1.DELETE("/stream/:id", handler.DeleteStream)
		v1.GET("/stream", handler.GetStreamList)
	}

	return nil
}
