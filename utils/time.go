package utils

import "time"

func GetDofTime() time.Time {
	// 获取当前时间
	currentTime := time.Now()
	// 格式化成 ISO 8601 格式
	formattedTime := currentTime.Format("2006-01-02T15:04:05-07:00")
	// 将格式化后的时间字符串解析为时间类型
	parsedTime, err := time.Parse("2006-01-02T15:04:05-07:00", formattedTime)
	if err != nil {
		return time.Time{}
	}
	return parsedTime
}
