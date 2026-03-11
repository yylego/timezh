package timezh

import "time"

// NT is the numeric-to-time conversion component
// Provides methods to create time objects from numeric components
//
// NT 是全局数字转时间转换器
// 提供从数字组件创建时间对象的方法
var NT = &nt把数字转化为时间{}

// nt把数字转化为时间 constructs time objects from numeric date/time components
// Creates time instances through combining date/time values
//
// nt把数字转化为时间 从数字日期/时间组件构建时间对象
// 通过组合年、月、日、时、分、秒值创建时间实例
type nt把数字转化为时间 struct{}

// New由当日时分秒得到时间 creates a time object at the current date with specified time values
// v几时: the time-value (0-23)
// v几分: minute (0-59)
// v几秒: second (0-59)
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns time object at the current date with specified time
//
// New由当日时分秒得到时间 创建今天日期、指定时分秒的时间对象
// v几时: 小时（0-23）
// v几分: 分钟（0-59）
// v几秒: 秒数（0-59）
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回今天日期、指定时分秒的时间对象
func (nt *nt把数字转化为时间) New由当日时分秒得到时间(v几时, v几分, v几秒 int, z时区位置 *time.Location) time.Time {
	now := time.Now().In(z时区位置)
	return time.Date(now.Year(), now.Month(), now.Day(), v几时, v几分, v几秒, 0, z时区位置)
}

// New由年月日得到零时时间 creates a time object at midnight on specified date
// v年份: the date-element (e.g., 2025)
// v月份: month (1-12)
// v几号: date of month (1-31)
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns time object at 00:00:00 on the specified date
//
// New由年月日得到零时时间 创建指定日期零时（午夜）的时间对象
// v年份: 年份（例如 2025）
// v月份: 月份（1-12）
// v几号: 日期（1-31）
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回指定日期 00:00:00 的时间对象
func (nt *nt把数字转化为时间) New由年月日得到零时时间(v年份, v月份, v几号 int, z时区位置 *time.Location) time.Time {
	return time.Date(v年份, time.Month(v月份), v几号, 0, 0, 0, 0, z时区位置)
}
