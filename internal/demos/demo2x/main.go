package main

import (
	"fmt"
	"time"

	"github.com/yylego/timezh"
)

func main() {
	// String to Time conversion (字符串转时间)
	// Parse standard time string (解析标准时间字符串)
	t1, err := timezh.ST.Get时间("2025-05-16 19:11:45", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed Time: %v\n", t1)

	// Parse date only (仅解析日期)
	t2, err := timezh.ST.Get通过日期得到时间("2025-05-16", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Date Time: %v\n", t2)

	// Number to Time conversion (数字转时间)
	// Create time from date components (从日期组件创建时间)
	t3 := timezh.NT.New由年月日得到零时时间(2025, 5, 16, timezh.TZ.CST)
	fmt.Printf("From Components: %v\n", t3)

	// Timestamp to Time conversion (时间戳转时间)
	// Convert Unix timestamp to time (转换 Unix 时间戳为时间)
	timestamp := time.Now().Unix()
	t4 := timezh.NS.Get秒数时间(timestamp)
	fmt.Printf("From Timestamp: %v\n", t4)

	// Time comparison (时间比较)
	// Compare time with time string (比较时间与时间字符串)
	now := time.Now()
	isBefore := timezh.CS.B时间在时分秒后面(now, "23:59:59")
	fmt.Printf("Is Before 23:59:59: %v\n", isBefore)
}
