package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_const "redis_stream/const"
	"redis_stream/redis"
)

func CreateStream(c *gin.Context) {
	req := CreateStreamReq{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	message := map[string]interface{}{
		"name": req.Name,
		"age":  req.Age,
	}

	err = redis.AddStream(_const.StreamName, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "success")
}
