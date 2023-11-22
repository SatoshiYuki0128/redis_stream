package handler

import "fmt"

type Response struct {
	Status int          `json:"status"`
	Data   interface{}  `json:"data,omitempty"`
	Error  *errorStruct `json:"error,omitempty"`
}

type errorStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type CreateStreamReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c *CreateStreamReq) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("empty name")
	}

	if c.Age <= 0 {
		return fmt.Errorf("invalid age")
	}

	return nil
}
