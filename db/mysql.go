package db

import (
	"fmt"
	"git.woa.com/sr-devops/console-server/consts"
	"git.woa.com/sr-devops/console-server/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectMysql : connect to database mysql using gorm
// gorm (GO ORM for SQL): http://gorm.io/docs/connecting_to_the_database.html
func MysqlInit() {
	if util.DefaultGetEnv("GO_ENV", "") != consts.PRODUCTION_ENV {
		util.LoadEnv()
	}
	dbType := util.DefaultGetEnv("DB_TYPE", "mysql")
	dbHost := util.DefaultGetEnv("DB_HOST", "localhost")
	dbPort := util.DefaultGetEnv("DB_PORT", "3306")
	dbName := util.DefaultGetEnv("DB_NAME", "release_management")
	dbUser := util.DefaultGetEnv("DB_USER", "root")
	dbPassword := util.DefaultGetEnv("DB_PASSWORD", "")

	mysqlArgs := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(dbType, mysqlArgs)
	MySql = db

	if err != nil {
		panic("初始化数据库失败")
	} else {
		fmt.Println("Connect database successfully")
	}
}
