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
		c.JSON(http.StatusBadRequest, Response{
			Status: http.StatusBadRequest,
			Error: &errorStruct{
				Code:    _const.BadRequest,
				Message: "c.ShouldBindJSON fail",
				Details: err.Error(),
			},
		})
		return
	}

	err = req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status: http.StatusBadRequest,
			Error: &errorStruct{
				Code:    _const.BadRequest,
				Message: "req.Validate fail",
				Details: err.Error(),
			},
		})
		return
	}

	message := map[string]interface{}{
		"name": req.Name,
		"age":  req.Age,
	}

	err = redis.AddStream(_const.StreamName, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status: _const.InternalError,
			Error: &errorStruct{
				Code:    _const.InternalError,
				Message: "redis.AddStream error",
				Details: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   message,
	})
}

func DeleteStream(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, Response{
			Status: http.StatusInternalServerError,
			Error: &errorStruct{
				Code:    _const.BadRequest,
				Message: "empty id",
			},
		})
		return
	}

	err := redis.DeleteStream(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status: http.StatusInternalServerError,
			Error: &errorStruct{
				Code:    _const.InternalError,
				Message: "redis.DeleteStream",
				Details: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   id,
	})
}

func GetStreamList(c *gin.Context) {
	messages, err := redis.GetStreamList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status: http.StatusInternalServerError,
			Error: &errorStruct{
				Code:    _const.InternalError,
				Message: "redis.GetStreamList",
				Details: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   messages,
	})
}
