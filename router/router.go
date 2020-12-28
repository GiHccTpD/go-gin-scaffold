package router

import (
	"git.woa.com/sr-devops/console-server/controllers/healthCheck"
	"git.woa.com/sr-devops/console-server/controllers/v1/login"
	"git.woa.com/sr-devops/console-server/controllers/v1/service"
	_ "git.woa.com/sr-devops/console-server/docs"
	"git.woa.com/sr-devops/console-server/middlewares/logger"
	"git.woa.com/sr-devops/console-server/util"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"fmt"
)

func InitRouter() *gin.Engine {
	if err := util.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
	}

	router := gin.New()
	router.Use(logger.LoggerMiddleware())

	// swagger api
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/healthCheck", healthCheck.ServiceCheck)
	router.GET("/MySQLCheck", healthCheck.MySQLCheck)

	apiVersionOne := router.Group("/api/v1/")
	apiVersionOne.POST("/login", login.Login)
	//apiVersionOne.Use(jwt.Jwt())
	apiVersionOne.POST("/service", service.CreateReleaseService)
	return router
}
