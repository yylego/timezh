package timezh

import (
	"time"

	"github.com/yylego/erero"
)

// CS is the time-string comparison component
// Provides methods to compare time objects with time strings
//
// CS 是全局时间字符串比较处理器
// 提供将时间对象与时间字符串进行比较的方法
var CS = &cs比较时间和字符串{}

// cs比较时间和字符串 compares time objects with time strings
// The time strings must be zero-padded "15:04:05", so comparing the strings gives the same result as comparing the times; a string in a different shape gives an error
//
// cs比较时间和字符串 将时间对象与时间字符串进行比较
// 时刻字符串必须是两位补零的 "15:04:05"，这样字符串的先后就是时间的先后；格式不对返回错误
type cs比较时间和字符串 struct{}

// Is时刻之间 checks if a time falls within a time-of-moment range (inclusive)
// t: the time object to check
// s起始时刻: start time string in format "15:04:05"
// e结束时刻: end time string in format "15:04:05"
// Returns true if time-of-moment is between start and end (inclusive)
//
// Is时刻之间 检查时间是否在时刻范围内（包含边界）
// t: 要检查的时间对象
// s起始时刻: "15:04:05" 格式的起始时刻字符串
// e结束时刻: "15:04:05" 格式的结束时刻字符串
// 返回 true 如果时刻在起始和结束之间（含边界）
func (cs *cs比较时间和字符串) Is时刻之间(t time.Time, s起始时刻, e结束时刻 string) (bool, error) {
	if !RS.Match时刻(s起始时刻) {
		return false, erero.Errorf("起始时刻格式错误: %s", s起始时刻)
	}
	if !RS.Match时刻(e结束时刻) {
		return false, erero.Errorf("结束时刻格式错误: %s", e结束时刻)
	}
	v := TS.T时刻.Get转为文本(t)
	return v >= s起始时刻 && v <= e结束时刻, nil
}

// Is晚于时刻 checks if a time is past a given time-of-moment
// t: the time object to check
// s时刻: time-of-moment string in format "15:04:05"
// Returns true if the time-of-moment is past the given time
//
// Is晚于时刻 检查时间是否晚于给定时刻
// t: 要检查的时间对象
// s时刻: "15:04:05" 格式的时刻字符串
// 返回 true 如果时刻晚于给定时刻
func (cs *cs比较时间和字符串) Is晚于时刻(t time.Time, s时刻 string) (bool, error) {
	if !RS.Match时刻(s时刻) {
		return false, erero.Errorf("时刻格式错误: %s", s时刻)
	}
	return TS.T时刻.Get转为文本(t) > s时刻, nil
}

// Is早于时刻 checks if a time is before a given time-of-moment
// t: the time object to check
// s时刻: time-of-moment string in format "15:04:05"
// Returns true if the time-of-moment is before the given time
//
// Is早于时刻 检查时间是否早于给定时刻
// t: 要检查的时间对象
// s时刻: "15:04:05" 格式的时刻字符串
// 返回 true 如果时刻早于给定时刻
func (cs *cs比较时间和字符串) Is早于时刻(t time.Time, s时刻 string) (bool, error) {
	if !RS.Match时刻(s时刻) {
		return false, erero.Errorf("时刻格式错误: %s", s时刻)
	}
	return TS.T时刻.Get转为文本(t) < s时刻, nil
}
