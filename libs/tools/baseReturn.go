package Tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type baseReturn struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewBaseReturn(code int, msg string, data interface{}) baseReturn {
	var baseretutn = baseReturn{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	return baseretutn
}

func Error(c *gin.Context, code int, msg string, data interface{}) {
	var baseretutn = baseReturn{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, baseretutn)
}

func Success(c *gin.Context, code int, msg string, data interface{}) {
	var baseretutn = baseReturn{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, baseretutn)
}
