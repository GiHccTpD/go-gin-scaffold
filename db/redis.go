package db

import (
	"fmt"
	"git.woa.com/sr-devops/console-server/consts"
	"git.woa.com/sr-devops/console-server/util"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

func RedisInit() {
	if util.DefaultGetEnv("GO_ENV", "") != consts.PRODUCTION_ENV {
		util.LoadEnv()
	}

	// 初始化redis数据库
	redisHost := util.DefaultGetEnv("REDIS_HOST", "localhost")
	redisPort := util.DefaultGetEnv("REDIS_PORT", "6379")
	redisPassword := util.DefaultGetEnv("REDIS_PASSWORD", "")
	redisDBStr := util.DefaultGetEnv("REDIS_DB", "0")
	redisDB, err := strconv.Atoi(redisDBStr)
	if err != nil {
		log.Fatal("redis数据库配置有误", redisDBStr)
	}
	client := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               fmt.Sprintf("%v:%v", redisHost, redisPort),
		Dialer:             nil,
		OnConnect:          nil,
		Password:           redisPassword,
		DB:                 redisDB,
		MaxRetries:         3,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	Redis = client
}
