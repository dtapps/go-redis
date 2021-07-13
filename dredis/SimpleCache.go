package dredis

import (
	"encoding/json"
	"time"
)

const (
	SerializerJson   = "json"
	SerializerString = "string"
)

type DBGttJsonFunc func() interface{}

type DBGttStringFunc func() string

// SimpleCache 缓存
type SimpleCache struct {
	Operation      *StringOperation // 操作类
	Expire         time.Duration    // 过去时间
	DBGetterJson   DBGttJsonFunc    // 缓存不存在的操作
	DBGetterString DBGttStringFunc  // 缓存不存在的操作
	Serializer     string           // 序列化方式
}

func NewSimpleCache(operation *StringOperation, expire time.Duration, serializer string) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire, Serializer: serializer}
}

// SetCache 设置缓存
func (c *SimpleCache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleCache) GetCache(key string) (ret interface{}) {
	if c.Serializer == SerializerJson {
		f := func() string {
			obj := c.DBGetterJson()
			b, err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return string(b)
		}
		ret = c.Operation.Get(key).UnwrapOrElse(f)
		c.SetCache(key, ret)
	} else if c.Serializer == SerializerString {
		f := func() string {
			return c.DBGetterString()
		}
		ret = c.Operation.Get(key).UnwrapOrElse(f)
		c.SetCache(key, ret)
	}
	return
}
