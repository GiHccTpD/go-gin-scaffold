package healthCheck

import (
	"git.woa.com/sr-devops/console-server/consts"
	"git.woa.com/sr-devops/console-server/db"
	"git.woa.com/sr-devops/console-server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 服务存活检测
// @Description 服务存活检测
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /healthCheck [get]
func ServiceCheck(c *gin.Context) {
	appG := util.Gin{C: c}
	appG.Response(http.StatusOK, consts.SUCCESS, nil)
}

// @Summary MySQL存活检测
// @Description MySQL存活检测
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"内部错误"}"
// @Router /MySQLCheck [get]
func MySQLCheck(c *gin.Context) {
	appG := util.Gin{C: c}

	err := db.MySql.DB().Ping()
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR, nil)
	} else {
		appG.Response(http.StatusOK, consts.SUCCESS, nil)
	}
}
