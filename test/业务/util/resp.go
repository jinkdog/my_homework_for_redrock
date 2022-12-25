package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status"` //状态码
	Info   string `json:"info"`   //返回信息
}

var OK = respTemplate{
	Status: 200,
	Info:   "Success",
}

var ParamError = respTemplate{
	Status: 400,
	Info:   "params error",
}

var InternalError = respTemplate{ //服务器错误
	Status: 500,
	Info:   "internal error",
}

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, OK) //此处OK替代gin.H{}输入的内容
}

func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
}

func RespInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalError)
}

func NomalErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
