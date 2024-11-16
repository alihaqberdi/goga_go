package response

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, res any, code ...int) {
	c.JSON(getCode(200, code), map[string]any{
		"res":    res,
		"status": true,
	})
}

func Fail(c *gin.Context, msg string, code ...int) {
	c.JSON(getCode(400, code), map[string]any{
		"status": false,
		"msg":    msg,
	})
}

func FailErr(c *gin.Context, err error, code ...int) {
	Fail(c, err.Error(), code...)
}

// others

func HasErr(c *gin.Context, err error, code ...int) bool {

	if err != nil {
		FailErr(c, err, code...)
		return true
	}

	return false
}

func Finish[R any](c *gin.Context, res R, err error, code ...int) {

	if err != nil {
		FailErr(c, err, code...)
		return
	}

	Success(c, res, code...)
}

func getCode(c int, code []int) int {
	if len(code) == 1 {
		return code[0]
	}

	return c
}
