package dredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	Rdb *redis.Client
)

func Version() string {
	return "v1.0.6"
}

// InitRedis 初始化连接 普通连接
func InitRedis(host string, port int, password string, db int) (err error) {
	dsn := fmt.Sprintf("%s:%v", host, port)
	fmt.Printf("【redis.普通】数据库配置 %s \n", dsn)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     dsn,      // address
		Password: password, // no password set
		DB:       db,       // use default DB
		PoolSize: 100,      // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	return err
}
