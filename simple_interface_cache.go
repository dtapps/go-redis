package dredis

import (
	"encoding/json"
	"time"
)

type DBGttInterfaceFunc func() interface{}

// SimpleInterfaceCache 缓存
type SimpleInterfaceCache struct {
	Operation *StringOperation   // 操作类
	Expire    time.Duration      // 过去时间
	DBGetter  DBGttInterfaceFunc // 缓存不存在的操作 DB
}

func NewSimpleInterfaceCache(operation *StringOperation, expire time.Duration) *SimpleInterfaceCache {
	return &SimpleInterfaceCache{Operation: operation, Expire: expire}
}

// SetCache 设置缓存
func (c *SimpleInterfaceCache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleInterfaceCache) GetCache(key string) (ret string) {
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
