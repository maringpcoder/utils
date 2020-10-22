package ts

import "time"

//根据传入的时间戳计算出该时间戳所在月开始的时间戳
func GetMonthBeginTimestamp(now int64) int64 {
	zone := time.FixedZone("CST", 8*3600)
	thisTime := time.Unix(now, 0).In(zone)
	return time.Date(thisTime.Year(), thisTime.Month(), 1, 0, 0, 0, 0, thisTime.Location()).Unix()
}
