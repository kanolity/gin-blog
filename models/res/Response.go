package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ListResp[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
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

func OKWithList(list any, count int64, c *gin.Context) {
	OKWithData(ListResp[any]{Count: count, List: list}, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OKWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}

func OKWithNothing(c *gin.Context) {
	Result(SUCCESS, map[string]any{}, "成功", c)
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
