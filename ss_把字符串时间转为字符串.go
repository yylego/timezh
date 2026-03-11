package timezh

import (
	"fmt"

	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// ss把字符串时间转为字符串 converts time strings between different formats
// Provides methods to transform and extract parts from time strings
//
// ss把字符串时间转为字符串 在不同格式间转换时间字符串
// 提供转换和提取时间字符串部分的方法
type ss把字符串时间转为字符串 struct{}

// SS is the string-to-string component
// Provides methods to transform time strings between various formats
//
// SS 是全局字符串转字符串转换器
// 提供在各种格式间转换时间字符串的方法
var SS = &ss把字符串时间转为字符串{}

// Get把年月日和时分秒分开返回 splits a datetime string into date and time parts
// s标准时间: datetime string in format "2006-01-02 15:04:05"
// Returns date part and time part, and/if blank strings if parsing fails
//
// Get把年月日和时分秒分开返回 将日期时间字符串分离为日期和时间部分
// s标准时间: "2006-01-02 15:04:05" 格式的日期时间字符串
// 返回日期部分和时间部分，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get把年月日和时分秒分开返回(s标准时间 string) (v标准日期, v标准时间 string) {
	if !RS.Match时间(s标准时间) {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间))
		return "", ""
	}
	num, err := fmt.Sscanf(s标准时间, "%s %s", &v标准日期, &v标准时间)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间), zap.Error(err))
		return "", ""
	}
	if num != 2 {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间), zap.Int("num", num))
		return "", ""
	}
	return v标准日期, v标准时间
}

// Get把8位数日期转换为标准日期 converts 8-digit date to standard format
// s八位日期: date string in format "20060102"
// Returns date string in format "2006-01-02", and/if blank if parsing fails
//
// Get把8位数日期转换为标准日期 将8位数日期转换为标准格式
// s八位日期: "20060102" 格式的日期字符串
// 返回 "2006-01-02" 格式的日期字符串，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get把8位数日期转换为标准日期(s八位日期 string) (s标准日期 string) {
	if !RS.Match日期8位数(s八位日期) {
		zaplog.LOG.Warn("日期格式错误", zap.String("s", s八位日期))
		return ""
	}
	return fmt.Sprintf("%s-%s-%s", s八位日期[0:4], s八位日期[4:6], s八位日期[6:8])
}

// Get把标准日期转换为8位数日期 converts standard date to 8-digit format
// s标准日期: date string in format "2006-01-02"
// Returns date string in format "20060102", and/if blank if parsing fails
//
// Get把标准日期转换为8位数日期 将标准日期转换为8位数格式
// s标准日期: "2006-01-02" 格式的日期字符串
// 返回 "20060102" 格式的日期字符串，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get把标准日期转换为8位数日期(s标准日期 string) (s八位日期 string) {
	if !RS.Match日期(s标准日期) {
		zaplog.LOG.Warn("日期格式错误", zap.String("s", s标准日期))
		return ""
	}
	return fmt.Sprintf("%s%s%s", s标准日期[0:4], s标准日期[5:7], s标准日期[8:10])
}

// Get把标准时间转换为8位数日期 extracts 8-digit date from datetime string
// s标准时间: datetime string in format "2006-01-02 15:04:05"
// Returns date part in format "20060102", and/if blank if parsing fails
//
// Get把标准时间转换为8位数日期 从日期时间字符串提取8位数日期
// s标准时间: "2006-01-02 15:04:05" 格式的日期时间字符串
// 返回 "20060102" 格式的日期部分，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get把标准时间转换为8位数日期(s标准时间 string) (s八位日期 string) {
	if !RS.Match时间(s标准时间) {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间))
		return ""
	}
	return fmt.Sprintf("%s%s%s", s标准时间[0:4], s标准时间[5:7], s标准时间[8:10])
}

// Get获取日期里的年份 extracts date-element from a date string
// s标准日期: date string in format "2006-01-02"
// Returns date-element string "2006", and/if blank if parsing fails
//
// Get获取日期里的年份 从日期字符串提取年份
// s标准日期: "2006-01-02" 格式的日期字符串
// 返回年份字符串 "2006"，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get获取日期里的年份(s标准日期 string) (v年份 string) {
	if !RS.Match日期(s标准日期) {
		zaplog.LOG.Warn("日期格式错误", zap.String("s", s标准日期))
		return ""
	}
	return s标准日期[0:4]
}

// Get获取日期中的月份和杠和几号 extracts month-date from a date string
// s标准日期: date string in format "2006-01-02"
// Returns month-date string "01-02", and/if blank if parsing fails
//
// Get获取日期中的月份和杠和几号 从日期字符串提取月-日
// s标准日期: "2006-01-02" 格式的日期字符串
// 返回月-日字符串 "01-02"，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get获取日期中的月份和杠和几号(s标准日期 string) (v月份和杠和几号 string) {
	if !RS.Match日期(s标准日期) {
		zaplog.LOG.Warn("日期格式错误", zap.String("s", s标准日期))
		return ""
	}
	return fmt.Sprintf("%s-%s", s标准日期[5:7], s标准日期[8:10])
}

// Get把6位数时刻转换为标准时刻 converts 6-digit time to standard format
// s六位时刻: time string in format "150405"
// Returns time string in format "15:04:05", and/if blank if parsing fails
//
// Get把6位数时刻转换为标准时刻 将6位数时刻转换为标准格式
// s六位时刻: "150405" 格式的时刻字符串
// 返回 "15:04:05" 格式的时刻字符串，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get把6位数时刻转换为标准时刻(s六位时刻 string) (s标准时刻 string) {
	if !RS.Match时刻6位数(s六位时刻) {
		zaplog.LOG.Warn("时刻格式错误", zap.String("s", s六位时刻))
		return ""
	}
	return fmt.Sprintf("%s:%s:%s", s六位时刻[0:2], s六位时刻[2:4], s六位时刻[4:6])
}

// Get把8位数带毫秒的时刻转换为标准时刻 converts 8-digit time (with milliseconds) to standard format
// s八位带毫秒时刻: time string in format "15040599" (8 digits with milliseconds)
// Returns time string in format "15:04:05", and/if blank if parsing fails
//
// Get把8位数带毫秒的时刻转换为标准时刻 将8位数带毫秒时刻转换为标准格式
// s八位带毫秒时刻: "15040599" 格式的8位数时刻字符串（含毫秒）
// 返回 "15:04:05" 格式的时刻字符串，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get把8位数带毫秒的时刻转换为标准时刻(s八位带毫秒时刻 string) (s标准时刻 string) {
	if !RS.Match时刻8位带毫秒纯数(s八位带毫秒时刻) {
		zaplog.LOG.Warn("时刻格式错误", zap.String("s", s八位带毫秒时刻))
		return ""
	}
	return fmt.Sprintf("%s:%s:%s", s八位带毫秒时刻[0:2], s八位带毫秒时刻[2:4], s八位带毫秒时刻[4:6])
}

// Get把标准时刻转换为6位数时刻 converts standard time to 6-digit format
// s标准时刻: time string in format "15:04:05"
// Returns time string in format "150405", and/if blank if parsing fails
//
// Get把标准时刻转换为6位数时刻 将标准时刻转换为6位数格式
// s标准时刻: "15:04:05" 格式的时刻字符串
// 返回 "150405" 格式的时刻字符串，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get把标准时刻转换为6位数时刻(s标准时刻 string) (s六位时刻 string) {
	if !RS.Match时刻(s标准时刻) {
		zaplog.LOG.Warn("时刻格式错误", zap.String("s", s标准时刻))
		return ""
	}
	return fmt.Sprintf("%s%s%s", s标准时刻[0:2], s标准时刻[3:5], s标准时刻[6:8])
}

// Get把4位时分转换为时和杠和分 formats 4-digit time-value-minute to standard format
// s四位时分数: time string in format "1504" (HHMM)
// Returns time string in format "15:04"
//
// Get把4位时分转换为时和杠和分 将4位时分格式化为标准格式
// s四位时分数: "1504" 格式的时分字符串（HHMM）
// 返回 "15:04" 格式的时间字符串
func (*ss把字符串时间转为字符串) Get把4位时分转换为时和杠和分(s四位时分数 string) (v时间和冒号和分数 string) {
	return fmt.Sprintf("%s:%s", s四位时分数[0:2], s四位时分数[2:4])
}

// Get通过标准时间得到日期 extracts date from a datetime string
// s: datetime string in format "2006-01-02 15:04:05"
// Returns date string in format "2006-01-02", and/if blank if parsing fails
//
// Get通过标准时间得到日期 从日期时间字符串提取日期
// s: "2006-01-02 15:04:05" 格式的日期时间字符串
// 返回 "2006-01-02" 格式的日期字符串，解析失败时返回空字符串
func (*ss把字符串时间转为字符串) Get通过标准时间得到日期(s string) (v标准日期 string) {
	if !RS.Match时间(s) {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s))
		return ""
	}
	return fmt.Sprintf("%s-%s-%s", s[0:4], s[5:7], s[8:10])
}
