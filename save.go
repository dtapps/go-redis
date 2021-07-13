package dredis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"time"
)

func SaveMap(key string, info string) *simplejson.Json {
	ctx := context.Background()
	jsonErrorStatus := Rdb.Set(ctx, DataGdBuildChart(key), info, 86400*time.Second).Err()
	if jsonErrorStatus != nil {
		fmt.Printf("保存 redis【%s】出错：%v\n", DataGdBuildChart(key), jsonErrorStatus)
	} else {
		fmt.Printf("保存 redis【%s】成功\n", DataGdBuildChart(key))
	}
	BytesConv, _ := json.Marshal(info)
	js, _ := simplejson.NewJson(BytesConv)
	return js
}
