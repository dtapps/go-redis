package dredis

import (
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	// 设置
	NewStringOperation().Set("name", "test", WithExpire(time.Second*20))

	// 获取
	iter := NewStringOperation().MGet("name", "age").Iter()
	for iter.HasNext() {
		log.Println(iter.Next())
	}

	// 假设缓存15s
	newCache := NewSimpleCache(NewStringOperation(), time.Second*15)
	// 缓存的key：news123 news101
	newCache.DBGetter = func() string {
		// 数据库获取
		return "data by id=123"
	}
	newCache.GetCache("news123")
}
