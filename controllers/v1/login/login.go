package login

import (
	"git.woa.com/sr-devops/console-server/consts"
	"git.woa.com/sr-devops/console-server/service/authentication"
	"git.woa.com/sr-devops/console-server/util"
	//"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	Username string `json:"Username" valid:"Required; MaxSize(50)"`
	Password string `json:"Password" valid:"Required; MaxSize(50)"`
}

// @Summary 登录
// @Description 登录
// @Accept  json
// @Produce  json
// @Param auth body Auth true "校验"
// @Success 200 {string} json "{"code":200,"data":{"token": "token"},"msg":"ok"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"TOKEN校验失败"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"TOKEN校验超时"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"生成TOKEN失败"}"
// @Failure 500 {string} json "{"code":500,"data":null,"msg":"错误的TOKEN"}"
// @Router /api/v1/login [post]
func Login(c *gin.Context) {
	appG := util.Gin{C: c}

	var json Auth
	if err := c.ShouldBindJSON(&json); err != nil {
		appG.Response(http.StatusBadRequest, consts.INVALID_PARAMS, nil)
		return
	}

	authService := authentication.Auth{Username: json.Username, Password: json.Password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusNotFound, consts.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(json.Username, json.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, consts.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
		"token": token,
	})
}
