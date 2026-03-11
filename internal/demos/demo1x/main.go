package main

import (
	"fmt"

	"github.com/yylego/timezh"
)

func main() {
	// Time to String conversion (时间转字符串)
	// Get current date in different formats (获取不同格式的当前日期)
	date := timezh.TS.D日期.Get现在日期(timezh.TZ.CST)
	fmt.Printf("Current Date: %s\n", date) // Output: 2025-05-16

	// Get current time (获取当前时间)
	s时间字符串 := timezh.TS.T时间.Get现在时间(timezh.TZ.CST)
	fmt.Printf("Current Time: %s\n", s时间字符串) // Output: 2025-05-16 19:11:45

	// Get yesterday's date (获取昨天的日期)
	yesterday := timezh.TS.D日期.Get昨日日期(timezh.TZ.CST)
	fmt.Printf("Yesterday: %s\n", yesterday) // Output: 2025-05-15

	// Get 8-digit date format (获取8位数字格式的日期)
	date8 := timezh.TS.D日期8位数.Get现在日期(timezh.TZ.CST)
	fmt.Printf("8-digit Date: %s\n", date8) // Output: 20250516

	// Get 14-digit time format (获取14位数字格式的时间)
	time14 := timezh.TS.T时间14位数.Get现在时间(timezh.TZ.CST)
	fmt.Printf("14-digit Time: %s\n", time14) // Output: 20250516191145
}
