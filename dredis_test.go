package dredis

import (
	"github.com/dtapps/go-redis/dredis"
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	// 连接
	err := dredis.InitRedis("redis", 6379, "", 2)
	if err != nil {
		panic(err)
	}

	// 设置
	dredis.NewStringOperation().Set("name", "test", dredis.WithExpire(time.Second*20))

	// 获取
	iter := dredis.NewStringOperation().MGet("name", "age").Iter()
	for iter.HasNext() {
		log.Println(iter.Next())
	}

	// 假设缓存15s
	newCache := dredis.NewSimpleCache(dredis.NewStringOperation(), time.Second*15, dredis.SerializerString)
	// 缓存的key：news123 news101
	newCache.DBGetter = func() interface{} {
		// 数据库获取
		return "data by id=123"
	}
	newCache.GetCache("news123")

}
