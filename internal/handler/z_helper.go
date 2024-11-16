package handler

import (
	"github.com/gin-gonic/gin"
)

func bind[D any](c *gin.Context) (*D, error) {
	data := new(D)
	err := c.Bind(data)

	return data, err
}
