package dredis

import (
	"context"
	"dtapps/dta/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	Rdb *redis.Client
)

// InitRedis 初始化连接
func InitRedis() (err error) {
	dsn := fmt.Sprintf("%s:%v", config.GlobConfig.RedisHost, config.GlobConfig.RedisPort)
	fmt.Printf("缓存数据库配置 %s \n", dsn)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: config.GlobConfig.RedisAuth, // no password set
		DB:       config.GlobConfig.RedisDB,   // use default DB
		PoolSize: 100,                         // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	return err
}
