package service

import (
	"git.woa.com/sr-devops/console-server/consts"
	"git.woa.com/sr-devops/console-server/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Service struct {
	WorkspaceID uint   `binding:"required,min=1"`
	ReleaseID   uint   `binding:"required,min=1"`
	ModuleID    uint   `binding:"required,min=1"`
	ServiceName string `binding:"required,min=1"`
}

// @Summary 创建发布服务
// @Description 创建发布服务
// @Accept  json
// @Produce  json
// @Param service body Service true "创建发布服务"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"创建失败"}"
// @Failure 500 {string} json "{"code":400,"data":null,"msg":{"WorkspaceID": "WorkspaceID为必填字段"}}"
// @Router /api/v1/service [post]
func CreateReleaseService(c *gin.Context) {
	appG := util.Gin{C: c}

	var service Service
	if err := c.ShouldBind(&service); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			appG.ResponseWithMsg(http.StatusBadRequest, consts.INVALID_PARAMS, nil, err.Error())
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code": consts.INVALID_PARAMS,
			"msg":  util.RemoveTopStruct(errs.Translate(util.Trans)),
			"data": nil,
		})

		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, nil)
}
