package handler

import (
	"fmt"
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

func DeleteStream(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, Response{
			http.StatusBadRequest,
			"empty id",
			nil,
		})
		return
	}

	err := redis.DeleteStream(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			http.StatusInternalServerError,
			err.Error(),
			nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		http.StatusOK,
		_const.SuccessMessage,
		nil,
	})
}

func GetStreamList(c *gin.Context) {
	messages, err := redis.GetStreamList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			http.StatusInternalServerError,
			fmt.Errorf("redis.GetStreamList err: %w", err).Error(),
			nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		http.StatusOK,
		"",
		messages,
	})
}
