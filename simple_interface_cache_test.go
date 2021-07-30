package dredis

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleInterfaceCache(t *testing.T) {
	newCache := NewSimpleInterfaceCache(NewStringOperation(), time.Second*10)
	newCache.DBGetter = func() interface{} {
		type a []string
		b := a{
			"me", "she", "you",
		}
		return b
	}
	fmt.Printf("TestSimpleInterfaceCache：%v\n", newCache.GetCache("TestSimpleInterfaceCache"))
}
