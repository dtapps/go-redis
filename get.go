package dredis

import (
	"context"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func GetMap(key string) (js *simplejson.Json) {
	ctx := context.Background()
	getInfo, err := Rdb.Get(ctx, DataGdBuildChart(key)).Result()
	if err != nil {
		fmt.Printf("获取 redis【%s】出错：%v\n", DataGdBuildChart(key), err)
		return nil
	} else {
		fmt.Printf("获取 redis【%s】成功\n", DataGdBuildChart(key))
		js, _ = simplejson.NewJson([]byte(getInfo))
		return js
	}
}
