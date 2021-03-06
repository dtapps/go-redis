package dredis

import (
	"encoding/json"
	"time"
)

type DBGttJsonFunc func() interface{}

// SimpleJsonCache 缓存
type SimpleJsonCache struct {
	Operation *StringOperation // 操作类
	Expire    time.Duration    // 过去时间
	DBGetter  DBGttJsonFunc    // 缓存不存在的操作 DB
}

func NewSimpleJsonCache(operation *StringOperation, expire time.Duration) *SimpleJsonCache {
	return &SimpleJsonCache{Operation: operation, Expire: expire}
}

// SetCache 设置缓存
func (c *SimpleJsonCache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleJsonCache) GetCache(key string) (ret interface{}) {
	f := func() string {
		obj := c.DBGetter()
		b, err := json.Marshal(obj)
		if err != nil {
			return ""
		}
		return string(b)
	}
	ret = c.Operation.Get(key).UnwrapOrElse(f)
	c.SetCache(key, ret)
	return
}
