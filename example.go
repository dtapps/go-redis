package dredis

import "gitee.com/dtapps/go-redis/dredis"

func main() {
	// 连接
	err := dredis.InitRedis("127.0.0.1", 6379, "", 2)
	if err != nil {
		panic(err)
	}
}
