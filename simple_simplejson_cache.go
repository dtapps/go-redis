package dredis

import (
	"github.com/bitly/go-simplejson"
	"time"
)

type DBGttSimpleJsonFunc func() *simplejson.Json

// SimpleSimpleJsonCache 缓存
type SimpleSimpleJsonCache struct {
	Operation *StringOperation    // 操作类
	Expire    time.Duration       // 过去时间
	DBGetter  DBGttSimpleJsonFunc // 缓存不存在的操作 DB
}

func NewSimpleSimpleJsonCache(operation *StringOperation, expire time.Duration) *SimpleSimpleJsonCache {
	return &SimpleSimpleJsonCache{Operation: operation, Expire: expire}
}

// SetCache 设置缓存
func (c *SimpleSimpleJsonCache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleSimpleJsonCache) GetCache(key string) (js *simplejson.Json) {
	f := func() string {
		obj := c.DBGetter()
		encode, err := obj.Encode()
		if err != nil {
			return ""
		}
		return string(encode)
	}
	ret := c.Operation.Get(key).UnwrapOrElse(f)
	c.SetCache(key, ret)
	js, _ = simplejson.NewJson([]byte(ret))
	return
}
