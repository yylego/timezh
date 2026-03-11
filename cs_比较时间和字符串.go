package timezh

import (
	"time"
)

// CS is the time-string comparison component
// Provides methods to compare time objects with time strings
//
// CS 是全局时间字符串比较处理器
// 提供将时间对象与时间字符串进行比较的方法
var CS = &cs比较时间和字符串{}

// cs比较时间和字符串 compares time objects with time strings
// Supports range checks and comparisons using time-of-moment strings
//
// cs比较时间和字符串 将时间对象与时间字符串进行比较
// 支持使用时刻字符串进行范围检查和前后比较
type cs比较时间和字符串 struct{}

// I在时分秒内含边界 checks if a time falls within a time-of-moment range (inclusive)
// t: the time object to check
// s起始时刻: start time string in format "15:04:05"
// e结束时刻: end time string in format "15:04:05"
// Returns true if time-of-moment is between start and end (inclusive)
//
// I在时分秒内含边界 检查时间是否在时刻范围内（包含边界）
// t: 要检查的时间对象
// s起始时刻: "15:04:05" 格式的起始时刻字符串
// e结束时刻: "15:04:05" 格式的结束时刻字符串
// 返回 true 如果时刻在起始和结束之间（含边界）
func (cs *cs比较时间和字符串) I在时分秒内含边界(t time.Time, s起始时刻, e结束时刻 string) bool {
	v := TS.T时刻.Get转字符串(t)
	return v >= s起始时刻 && v <= e结束时刻
}

// A时间在时分秒前面 checks if a time is past a given time-of-moment
// t: the time object to check
// s时刻: time-of-moment string in format "15:04:05"
// Returns true if the time-of-moment is past the given time
//
// A时间在时分秒前面 检查时间是否晚于给定时刻
// t: 要检查的时间对象
// s时刻: "15:04:05" 格式的时刻字符串
// 返回 true 如果时刻晚于给定时刻
func (cs *cs比较时间和字符串) A时间在时分秒前面(t time.Time, s时刻 string) bool {
	return TS.T时刻.Get转字符串(t) > s时刻
}

// B时间在时分秒后面 checks if a time is before a given time-of-moment
// t: the time object to check
// s时刻: time-of-moment string in format "15:04:05"
// Returns true if the time-of-moment is before the given time
//
// B时间在时分秒后面 检查时间是否早于给定时刻
// t: 要检查的时间对象
// s时刻: "15:04:05" 格式的时刻字符串
// 返回 true 如果时刻早于给定时刻
func (cs *cs比较时间和字符串) B时间在时分秒后面(t time.Time, s时刻 string) bool {
	return TS.T时刻.Get转字符串(t) < s时刻
}
