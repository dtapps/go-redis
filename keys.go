package dredis

import "fmt"

// DataGdBuildChart 数据分析 广东省项目 视图
func DataGdBuildChart(id string) string {
	return fmt.Sprintf("data:gdbuild:chart:%s", id)
}

func ApiIpQqWry(id string) string {
	return fmt.Sprintf("api:ip:qqwry:%s", id)
}
