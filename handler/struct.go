package handler

import "fmt"

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
