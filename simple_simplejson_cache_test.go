package dredis

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"testing"
	"time"
)

func TestSimpleSimpleJsonCache(t *testing.T) {
	newCache := NewSimpleSimpleJsonCache(NewStringOperation(), time.Second*10)
	newCache.DBGetter = func() *simplejson.Json {
		a := simplejson.New()
		a.Set("me", "you")
		return a
	}
	fmt.Printf("TestSimpleSimpleJsonCache：%v\n", newCache.GetCache("TestSimpleSimpleJsonCache"))
}
