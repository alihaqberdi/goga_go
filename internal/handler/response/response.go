package response

import (
	"errors"
	"github.com/alihaqberdi/goga_go/internal/pkg/app_errors"
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, res any, code ...int) {
	//c.JSON(getCode(200, code), map[string]any{
	//	"res":    res,
	//	"status": true,
	//})
	c.JSON(getCode(200, code), res)
}

func Fail(c *gin.Context, msg string, code ...int) {
	c.JSON(getCode(400, code), map[string]any{
		"status":  false,
		"message": msg,
	})
}

func FailErr(c *gin.Context, err error, code ...int) {

	var appErr *app_errors.AppError
	if errors.As(err, &appErr) {
		Fail(c, appErr.Message, appErr.Status)
		return
	}

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
