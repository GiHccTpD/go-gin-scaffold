package db

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var MySql *gorm.DB

var Redis *redis.Client

func init() {
	// 初始化mysql
	MysqlInit()
	// 初始化redis
	RedisInit()
}
