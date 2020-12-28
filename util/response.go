package util

import (
	"git.woa.com/sr-devops/console-server/consts"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  consts.GetMsg(errCode),
		"data": data,
	})

	return
}

func (g *Gin) ResponseWithMsg(httpCode, errCode int, data interface{}, msg string) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  msg,
		"data": data,
	})

	return
}
