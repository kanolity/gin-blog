package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = -1
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func OK(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}
func OKWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}
func OKWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}

func Fail(data interface{}, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}
func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(ERROR, map[string]any{}, "未知错误", c)
}
