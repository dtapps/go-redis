package dredis

import (
	"github.com/bitly/go-simplejson"
	"github.com/dtapps/go-redis/dredis"
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	log.Printf("version：%v\n", dredis.Version())
	// 连接
	//err := dredis.InitRedis("127.0.0.1", 6379, "", 2)
	//if err != nil {
	//	panic(err)
	//}
	//jsonSimpleJson()
}

func set() {
	// 设置
	dredis.NewStringOperation().Set("test", "test", dredis.WithExpire(time.Second*1))
}

func mGet() {
	// 获取
	iter := dredis.NewStringOperation().MGet("test1", "test2").Iter()
	for iter.HasNext() {
		log.Println("MGet：", iter.Next())
	}
}

func json() {
	newCache := dredis.NewSimpleCache(dredis.NewStringOperation(), time.Second*10, dredis.SerializerJson)
	newCache.JsonGetter = func() interface{} {
		log.Println("【没有命中】SerializerJson")
		type a []string
		b := a{
			"me", "she", "you",
		}
		return b
	}
	cacheJSon := newCache.GetCache("test123")
	log.Printf("【GetCache】cacheJSon：%v\n", cacheJSon)
}

func dbString() {
	newCache := dredis.NewSimpleCache(dredis.NewStringOperation(), time.Second*10, dredis.SerializerString)
	newCache.DBGetter = func() string {
		log.Println("【没有命中】SerializerString")
		return "data by id=123"
	}
	cacheString := newCache.GetCache("test456")
	log.Printf("【GetCache】cacheString：%v\n", cacheString)
}

func simpleJson() {
	newCache := dredis.NewSimpleCache(dredis.NewStringOperation(), time.Second*50, dredis.SerializerSimpleJson)
	newCache.SimpleJsonGetter = func() *simplejson.Json {
		log.Println("_test【没有命中】SerializerSimpleJson")
		js := simplejson.New()
		js.Set("name", "test")
		return js
	}
	cacheSimpleJson := newCache.GetCacheSimpleJson("test789")
	log.Printf("_test【GetCache】cacheSimpleJson：%v\n", cacheSimpleJson.Get("name"))
}

func jsonSimpleJson() {
	newCache := dredis.NewSimpleCache(dredis.NewStringOperation(), time.Second*50, dredis.SerializerJson)
	newCache.JsonGetter = func() interface{} {
		log.Println("【没有命中】SerializerJson")
		type a []string
		b := a{
			"me", "she", "you",
		}
		return b
	}
	cacheJson := newCache.GetCacheSimpleJson("test789")
	log.Printf("_test【JsonGetter GetCacheSimpleJson】jsonSimpleJson：%v\n", cacheJson.GetIndex(1))
}
