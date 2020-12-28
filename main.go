package main

import (
	"fmt"
	"git.woa.com/sr-devops/console-server/db"
	"git.woa.com/sr-devops/console-server/models"
	"git.woa.com/sr-devops/console-server/router"
	"git.woa.com/sr-devops/console-server/util"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {
	// 通过数据库管理gorm自动生成表
	db.MySql.AutoMigrate(models.AutoMigrateModels...)

	r := router.InitRouter()

	address := fmt.Sprintf("%s:%s", util.DefaultGetEnv("HOST", "localhost"), util.DefaultGetEnv("PORT", "10080"))
	server := endless.NewServer(address, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
