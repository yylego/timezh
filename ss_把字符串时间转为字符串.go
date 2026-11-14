package timezh

import (
	"fmt"

	"github.com/yylego/erero"
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

// Get拆分日期时刻 splits a datetime string into the date part and the clock time part
// s标准时间: datetime string in format "2006-01-02 15:04:05"
//
// Get拆分日期时刻 将日期时间字符串拆成日期和时刻两部分
// s标准时间: "2006-01-02 15:04:05" 格式的日期时间字符串
func (*ss把字符串时间转为字符串) Get拆分日期时刻(s标准时间 string) (v标准日期, v标准时刻 string, err error) {
	if !RS.Match时间(s标准时间) {
		return "", "", erero.Errorf("时间格式错误: %s", s标准时间)
	}
	return s标准时间[0:10], s标准时间[11:19], nil
}

// Get标准日期 converts 8-digit date to standard format
// s八位日期: date string in format "20060102"
// Returns date string in format "2006-01-02"
//
// Get标准日期 将8位数日期转换为标准格式
// s八位日期: "20060102" 格式的日期字符串
// 返回 "2006-01-02" 格式的日期字符串
func (*ss把字符串时间转为字符串) Get标准日期(s八位日期 string) (string, error) {
	if !RS.Match八位数字日期(s八位日期) {
		return "", erero.Errorf("日期格式错误: %s", s八位日期)
	}
	return fmt.Sprintf("%s-%s-%s", s八位日期[0:4], s八位日期[4:6], s八位日期[6:8]), nil
}

// Get八位数字日期 converts standard date to 8-digit format
// s标准日期: date string in format "2006-01-02"
// Returns date string in format "20060102"
//
// Get八位数字日期 将标准日期转换为8位数格式
// s标准日期: "2006-01-02" 格式的日期字符串
// 返回 "20060102" 格式的日期字符串
func (*ss把字符串时间转为字符串) Get八位数字日期(s标准日期 string) (string, error) {
	if !RS.Match日期(s标准日期) {
		return "", erero.Errorf("日期格式错误: %s", s标准日期)
	}
	return fmt.Sprintf("%s%s%s", s标准日期[0:4], s标准日期[5:7], s标准日期[8:10]), nil
}

// Get年份 extracts date-element from a date string
// s标准日期: date string in format "2006-01-02"
// Returns date-element string "2006"
//
// Get年份 从日期字符串提取年份
// s标准日期: "2006-01-02" 格式的日期字符串
// 返回年份字符串 "2006"
func (*ss把字符串时间转为字符串) Get年份(s标准日期 string) (string, error) {
	if !RS.Match日期(s标准日期) {
		return "", erero.Errorf("日期格式错误: %s", s标准日期)
	}
	return s标准日期[0:4], nil
}

// Get日期月日 extracts month-date from a date string
// s标准日期: date string in format "2006-01-02"
// Returns month-date string "01-02"
//
// Get日期月日 从日期字符串提取月-日
// s标准日期: "2006-01-02" 格式的日期字符串
// 返回月-日字符串 "01-02"
func (*ss把字符串时间转为字符串) Get日期月日(s标准日期 string) (string, error) {
	if !RS.Match日期(s标准日期) {
		return "", erero.Errorf("日期格式错误: %s", s标准日期)
	}
	return fmt.Sprintf("%s-%s", s标准日期[5:7], s标准日期[8:10]), nil
}

// Get标准时刻 converts 6-digit time to standard format
// s六位时刻: time string in format "150405"
// Returns time string in format "15:04:05"
//
// Get标准时刻 将6位数时刻转换为标准格式
// s六位时刻: "150405" 格式的时刻字符串
// 返回 "15:04:05" 格式的时刻字符串
func (*ss把字符串时间转为字符串) Get标准时刻(s六位时刻 string) (string, error) {
	if !RS.Match六位数字时刻(s六位时刻) {
		return "", erero.Errorf("时刻格式错误: %s", s六位时刻)
	}
	return fmt.Sprintf("%s:%s:%s", s六位时刻[0:2], s六位时刻[2:4], s六位时刻[4:6]), nil
}

// Get八位数字转标准时刻 converts 8-digit time to standard format, dropping the 2 digits of the second's fraction
// s八位时刻: time string in format "15040599", "150405" and then 2 digits below the second
// Returns time string in format "15:04:05"
//
// Get八位数字转标准时刻 将8位数时刻转换为标准格式，去掉末尾两位秒以下的小数
// s八位时刻: "15040599" 格式的时刻字符串，"150405" 后面跟两位秒以下的小数
// 返回 "15:04:05" 格式的时刻字符串
func (*ss把字符串时间转为字符串) Get八位数字转标准时刻(s八位时刻 string) (string, error) {
	if !RS.Match八位数字时刻(s八位时刻) {
		return "", erero.Errorf("时刻格式错误: %s", s八位时刻)
	}
	return fmt.Sprintf("%s:%s:%s", s八位时刻[0:2], s八位时刻[2:4], s八位时刻[4:6]), nil
}

// Get六位数字时刻 converts standard time to 6-digit format
// s标准时刻: time string in format "15:04:05"
// Returns time string in format "150405"
//
// Get六位数字时刻 将标准时刻转换为6位数格式
// s标准时刻: "15:04:05" 格式的时刻字符串
// 返回 "150405" 格式的时刻字符串
func (*ss把字符串时间转为字符串) Get六位数字时刻(s标准时刻 string) (string, error) {
	if !RS.Match时刻(s标准时刻) {
		return "", erero.Errorf("时刻格式错误: %s", s标准时刻)
	}
	return fmt.Sprintf("%s%s%s", s标准时刻[0:2], s标准时刻[3:5], s标准时刻[6:8]), nil
}

// Get标准时分 converts the 4-digit "1504" form to standard format
// s四位时分数: time string in format "1504" (HHMM)
// Returns time string in format "15:04"
//
// Get标准时分 将4位时分转换为标准格式
// s四位时分数: "1504" 格式的时分字符串（HHMM）
// 返回 "15:04" 格式的时分字符串
func (*ss把字符串时间转为字符串) Get标准时分(s四位时分数 string) (string, error) {
	if !RS.Match四位数字时分(s四位时分数) {
		return "", erero.Errorf("时分格式错误: %s", s四位时分数)
	}
	return fmt.Sprintf("%s:%s", s四位时分数[0:2], s四位时分数[2:4]), nil
}

// Get取出日期 extracts date from a datetime string
// s: datetime string in format "2006-01-02 15:04:05"
// Returns date string in format "2006-01-02"
//
// Get取出日期 从日期时间字符串提取日期
// s: "2006-01-02 15:04:05" 格式的日期时间字符串
// 返回 "2006-01-02" 格式的日期字符串
func (*ss把字符串时间转为字符串) Get取出日期(s string) (string, error) {
	if !RS.Match时间(s) {
		return "", erero.Errorf("时间格式错误: %s", s)
	}
	return fmt.Sprintf("%s-%s-%s", s[0:4], s[5:7], s[8:10]), nil
}

// Get取出时刻 extracts the clock time from a datetime string
// s: datetime string in format "2006-01-02 15:04:05"
// Returns time-of-moment string in format "15:04:05"
//
// Get取出时刻 从日期时间字符串提取时刻
// s: "2006-01-02 15:04:05" 格式的日期时间字符串
// 返回 "15:04:05" 格式的时刻字符串
func (*ss把字符串时间转为字符串) Get取出时刻(s string) (string, error) {
	if !RS.Match时间(s) {
		return "", erero.Errorf("时间格式错误: %s", s)
	}
	return s[11:19], nil
}
