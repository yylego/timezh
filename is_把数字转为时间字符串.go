package timezh

import "fmt"

// IS is the numeric-to-string conversion component
// Provides methods to format numeric date/time components into time strings
//
// IS 是全局整数转字符串转换器
// 提供将数字日期/时间组件格式化为时间字符串的方法
var IS = is把数字转为时间字符串{}

// is把数字转为时间字符串 formats numeric components into time strings
// Converts date/time numeric values into formatted strings
//
// is把数字转为时间字符串 将数字组件格式化为时间字符串
// 将年、月、日、时、分、秒数值转换为格式化字符串
type is把数字转为时间字符串 struct{}

// Get日期 formats date components into a date string
// v年份: the given date-element (e.g., 2025)
// v月份: month (1-12)
// v几号: date of month (1-31)
// Returns formatted date string "2006-01-02"
//
// Get日期 将年、月、日格式化为日期字符串
// v年份: 年份（例如 2025）
// v月份: 月份（1-12）
// v几号: 日期（1-31）
// 返回格式化的日期字符串 "2006-01-02"
func (si *is把数字转为时间字符串) Get日期(v年份, v月份, v几号 int) string {
	return fmt.Sprintf("%04d-%02d-%02d", v年份, v月份, v几号)
}

// Get时刻 formats time-of-moment components into a string
// v几时: the time-of-moment (0-23)
// v几分: minute (0-59)
// v几秒: second (0-59)
// Returns formatted time string "15:04:05"
//
// Get时刻 将时、分、秒格式化为时刻字符串
// v几时: 小时（0-23）
// v几分: 分钟（0-59）
// v几秒: 秒数（0-59）
// 返回格式化的时刻字符串 "15:04:05"
func (si *is把数字转为时间字符串) Get时刻(v几时, v几分, v几秒 int) string {
	return fmt.Sprintf("%02d:%02d:%02d", v几时, v几分, v几秒)
}

// Get时间 formats complete date/time components into a datetime string
// v年份: the given date-element (e.g., 2025)
// v月份: month (1-12)
// v几号: date of month (1-31)
// v几时: the time-of-moment (0-23)
// v几分: minute (0-59)
// v几秒: second (0-59)
// Returns formatted datetime string "2006-01-02 15:04:05"
//
// Get时间 将所有日期/时间组件格式化为日期时间字符串
// v年份: 年份（例如 2025）
// v月份: 月份（1-12）
// v几号: 日期（1-31）
// v几时: 小时（0-23）
// v几分: 分钟（0-59）
// v几秒: 秒数（0-59）
// 返回格式化的日期时间字符串 "2006-01-02 15:04:05"
func (si *is把数字转为时间字符串) Get时间(v年份, v月份, v几号, v几时, v几分, v几秒 int) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", v年份, v月份, v几号, v几时, v几分, v几秒)
}
