package dredis

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleStringCache(t *testing.T) {
	newCache := NewSimpleStringCache(NewStringOperation(), time.Second*10)
	newCache.DBGetter = func() string {
		return "me"
	}
	fmt.Printf("TestSimpleStringCache：%v\n", newCache.GetCache("TestSimpleStringCache"))
}
