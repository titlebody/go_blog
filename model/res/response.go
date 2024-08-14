package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	Success = 0
	Err     = 7
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data interface{}, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func Fail(data interface{}, msg string, c *gin.Context) {
	Result(Err, data, msg, c)
}

func FailWithMessage(data interface{}, c *gin.Context) {
	Result(Err, data, "失败", c)
}

func FailWithCode(code int, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(Err, map[string]any{}, msg, c)
		return
	}
	Result(Err, map[string]any{}, "未知错误", c)
}
