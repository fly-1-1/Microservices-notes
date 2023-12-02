package models

import (
	"time"
)

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func GetDay() string {
	// 获取当前时间
	currentTime := time.Now()

	// 使用Format方法将时间格式化为"20231203"的格式
	formattedDate := currentTime.Format("20060102")

	return formattedDate
}

func GetTimeStamp() string {
	// 获取当前时间的纳秒级时间戳
	currentTimestamp := time.Now().UnixNano()

	// 将纳秒转换为毫秒
	currentTimestampInMilliseconds := currentTimestamp / int64(time.Millisecond)

	// 将毫秒级时间戳转换为时间对象
	currentTime := time.Unix(0, currentTimestampInMilliseconds*int64(time.Millisecond))

	// 使用Format方法将时间格式化为所需的日期时间格式，包含毫秒
	formattedTime := currentTime.Format("20060102150405")

	return formattedTime
}
