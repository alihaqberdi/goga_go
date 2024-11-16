package response

import (
	"github.com/alihaqberdi/goga_go/internal/pkg/status"
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, res interface{}, code ...status.Code) {
	c.JSON(getCode(200, code), map[string]interface{}{
		"res":    res,
		"status": true,
	})
}

func Fail(c *gin.Context, msg string, code ...status.Code) {
	c.JSON(getCode(400, code), map[string]interface{}{
		"status": false,
		"msg":    msg,
	})
}

func FailErr(c *gin.Context, err error, code ...status.Code) {
	Fail(c, err.Error(), code...)
}

// others

func HasErr(c *gin.Context, err error, code ...status.Code) bool {

	if err != nil {
		FailErr(c, err, code...)
		return true
	}

	return false
}

func Finish[R any](c *gin.Context, res R, err error, code ...status.Code) {

	if err != nil {
		FailErr(c, err, code...)
		return
	}

	Success(c, res, code...)
}

func getCode(c int, code []status.Code) int {
	if len(code) == 1 {
		return code[0].GetOrDef(c)
	}

	return c
}
