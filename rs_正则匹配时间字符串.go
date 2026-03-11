package timezh

import "regexp"

// RS is the regex-based time string component
// Provides compiled regex patterns and methods to validate time string formats
//
// RS 是全局基于正则的时间字符串匹配器
// 提供编译好的正则表达式模式和验证时间字符串格式的方法
var RS = &rs正则匹配时间字符串{
	regexp日期:          regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
	regexp时刻:          regexp.MustCompile(`^\d{2}:\d{2}:\d{2}$`),
	regexp时间:          regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`),
	regexp日期和零时零分零秒时间: regexp.MustCompile(`^\d{4}-\d{2}-\d{2} 00:00:00$`),
	regexp日期8位数:       regexp.MustCompile(`^\d{8}$`),
	regexp时刻6位数:       regexp.MustCompile(`^\d{6}$`),
	regexp时刻8位带毫秒纯数:   regexp.MustCompile(`^\d{8}$`),
	regexp年份:          regexp.MustCompile(`^\d{4}$`),
}

// rs正则匹配时间字符串 validates time strings using regex patterns
// Contains compiled patterns on various date/time formats
//
// rs正则匹配时间字符串 使用正则表达式模式验证时间字符串
// 包含各种日期/时间格式的编译模式
type rs正则匹配时间字符串 struct {
	regexp日期          *regexp.Regexp // Pattern for "2006-01-02" | "2006-01-02" 格式
	regexp时刻          *regexp.Regexp // Pattern for "15:04:05" | "15:04:05" 格式
	regexp时间          *regexp.Regexp // Pattern for "2006-01-02 15:04:05" | "2006-01-02 15:04:05" 格式
	regexp日期和零时零分零秒时间 *regexp.Regexp // Pattern for "2006-01-02 00:00:00" | "2006-01-02 00:00:00" 格式
	regexp日期8位数       *regexp.Regexp // Pattern for "20060102" | "20060102" 格式
	regexp时刻6位数       *regexp.Regexp // Pattern for "150405" | "150405" 格式
	regexp时刻8位带毫秒纯数   *regexp.Regexp // Pattern for 8-digit time with milliseconds | 8位数带毫秒时刻
	regexp年份          *regexp.Regexp // Pattern for "2006" | "2006" 格式
}

// Match日期 validates if string matches date format "2006-01-02"
// s标准日期: string to validate
// Returns true if format matches
//
// Match日期 验证字符串是否匹配日期格式 "2006-01-02"
// s标准日期: 要验证的字符串
// 返回 true 如果格式匹配
func (match *rs正则匹配时间字符串) Match日期(s标准日期 string) bool {
	return match.regexp日期.MatchString(s标准日期)
}

// Match时刻 validates if string matches time format "15:04:05"
// s标准时刻: string to validate
// Returns true if format matches
//
// Match时刻 验证字符串是否匹配时刻格式 "15:04:05"
// s标准时刻: 要验证的字符串
// 返回 true 如果格式匹配
func (match *rs正则匹配时间字符串) Match时刻(s标准时刻 string) bool {
	return match.regexp时刻.MatchString(s标准时刻)
}

// Match时间 validates if string matches datetime format "2006-01-02 15:04:05"
// s标准时间: string to validate
// Returns true if format matches
//
// Match时间 验证字符串是否匹配日期时间格式 "2006-01-02 15:04:05"
// s标准时间: 要验证的字符串
// 返回 true 如果格式匹配
func (match *rs正则匹配时间字符串) Match时间(s标准时间 string) bool {
	return match.regexp时间.MatchString(s标准时间)
}

// Match日期和零时零分零秒时间 validates if string matches midnight format "2006-01-02 00:00:00"
// s: string to validate
// Returns true if format matches
//
// Match日期和零时零分零秒时间 验证字符串是否匹配零时格式 "2006-01-02 00:00:00"
// s: 要验证的字符串
// 返回 true 如果格式匹配
func (match *rs正则匹配时间字符串) Match日期和零时零分零秒时间(s string) bool {
	return match.regexp日期和零时零分零秒时间.MatchString(s)
}

// Match日期8位数 validates if string matches 8-digit date format "20060102"
// s: string to validate
// Returns true if format matches
//
// Match日期8位数 验证字符串是否匹配8位数日期格式 "20060102"
// s: 要验证的字符串
// 返回 true 如果格式匹配
func (match *rs正则匹配时间字符串) Match日期8位数(s string) bool {
	return match.regexp日期8位数.MatchString(s)
}

// Match时刻6位数 validates if string matches 6-digit time format "150405"
// s: string to validate
// Returns true if format matches
//
// Match时刻6位数 验证字符串是否匹配6位数时刻格式 "150405"
// s: 要验证的字符串
// 返回 true 如果格式匹配
func (match *rs正则匹配时间字符串) Match时刻6位数(s string) bool {
	return match.regexp时刻6位数.MatchString(s)
}

// Match时刻8位带毫秒纯数 validates if string matches 8-digit time with milliseconds
// s: string to validate
// Returns true if format matches
//
// Match时刻8位带毫秒纯数 验证字符串是否匹配8位数带毫秒时刻
// s: 要验证的字符串
// 返回 true 如果格式匹配
func (match *rs正则匹配时间字符串) Match时刻8位带毫秒纯数(s string) bool {
	return match.regexp时刻8位带毫秒纯数.MatchString(s)
}

// Match年份 validates if string matches 4-digit date-element format "2006"
// s四位年份: string to validate
// Returns true if format matches
//
// Match年份 验证字符串是否匹配4位数年份格式 "2006"
// s四位年份: 要验证的字符串
// 返回 true 如果格式匹配
func (match *rs正则匹配时间字符串) Match年份(s四位年份 string) bool {
	return match.regexp年份.MatchString(s四位年份)
}
