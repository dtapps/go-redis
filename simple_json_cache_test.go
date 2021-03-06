package dredis

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleJsonCache(t *testing.T) {
	newCache := NewSimpleJsonCache(NewStringOperation(), time.Second*10)
	newCache.DBGetter = func() interface{} {
		type a []string
		b := a{
			"me", "she", "you",
		}
		return b
	}
	fmt.Printf("TestSimpleJsonCacheļ¼%v\n", newCache.GetCache("TestSimpleJsonCache"))
}
