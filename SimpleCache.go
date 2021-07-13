package dredis

import "time"

type DBGttFunc func() string

type SimpleCache struct {
	Operation *StringOperation
	Expire    time.Duration
	DBGetter  DBGttFunc
}

func NewSimpleCache(operation *StringOperation, expire time.Duration) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire}
}

// SetCache 设置缓存
func (c *SimpleCache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleCache) GetCache(key string) (ret interface{}) {
	ret = c.Operation.Get(key).UnwrapOrElse(c.DBGetter)
	return
}
