package main

import (
	"fmt"

	"github.com/yylego/timezh"
)

func main() {
	// Working with Time Zones (时区操作)
	// Get current date in CST (获取 CST 时区的当前日期)
	dateCST := timezh.TS.D日期.Get当日日期(timezh.TZ.CST)
	fmt.Printf("CST Date: %s\n", dateCST) // Output: 2025-05-16

	// Get current date in UTC (获取 UTC 时区的当前日期)
	dateUTC := timezh.TS.D日期.Get当日日期(timezh.TZ.UTC)
	fmt.Printf("UTC Date: %s\n", dateUTC) // Output: 2025-05-16

	// Get current date in ICT (获取 ICT 时区的当前日期)
	dateICT := timezh.TS.D日期.Get当日日期(timezh.TZ.ICT)
	fmt.Printf("ICT Date: %s\n", dateICT) // Output: 2025-05-16

	// Get recent business days (获取最近的工作日)
	// Get the latest 5 Mon-Fri dates up to now (获取包括今天在内的最近5个工作日)
	workdays := timezh.TS.D日期.Get截至当日工作日期(5, timezh.TZ.CST)
	fmt.Printf("Recent Workdays: %v\n", workdays) // Output: [2025-05-16 2025-05-15 2025-05-14 2025-05-13 2025-05-12]

	// Calculate time difference (计算时间差)
	duration, err := timezh.SD.Get两个时间间隔("2025-05-16 19:00:00", "2025-05-16 20:00:00")
	if err != nil {
		panic(err)
	}
	hours := duration.Hours()
	fmt.Printf("Time Difference: %.0f hours\n", hours) // Output: Time Difference: 1 hours
}
