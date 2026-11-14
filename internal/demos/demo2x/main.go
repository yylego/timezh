package main

import (
	"fmt"
	"time"

	"github.com/yylego/timezh"
)

func main() {
	// String to Time conversion (字符串转时间)
	// Parse standard time string (解析标准时间字符串)
	t1, err := timezh.TS.T时间.Get转为时间("2025-05-16 19:11:45", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed Time: %v\n", t1)

	// Parse just date (仅解析日期)
	t2, err := timezh.TS.D日期.Get转为时间("2025-05-16", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Date Time: %v\n", t2)

	// Numeric to Time conversion (数字转时间)
	// Create time from date components (从日期组件创建时间)
	t3, err := timezh.NT.New某日零点时间(2025, 5, 16, timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("From Components: %v\n", t3)

	// Timestamp to Time conversion (时间戳转时间)
	// Convert Unix timestamp to time (转换 Unix 时间戳为时间)
	timestamp := time.Now().Unix()
	t4 := timezh.NS.Get秒数时间(timestamp, timezh.TZ.CST)
	fmt.Printf("From Timestamp: %v\n", t4)

	// Time comparison (时间比较)
	// Compare time and time string (比较时间与时间字符串)
	now := time.Now()
	isBefore, err := timezh.CS.Is早于时刻(now, "23:59:59")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Before 23:59:59: %v\n", isBefore)
}
