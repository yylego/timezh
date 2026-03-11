package timezh

import (
	"time"
)

// TS is the time-to-string component with predefined format templates
// Provides various date/time formatting options through structured fields
//
// TS 是全局时间转字符串转换器，包含预定义的格式模板
// 通过结构化字段提供各种日期/时间格式化选项
var TS = NewTs转字符串()

// ts转字符串 holds various time-to-string components with different layouts
// Each field provides specific date/time format conversion
//
// ts转字符串 包含各种不同布局的时间转字符串格式化器
// 每个字段提供特定的日期/时间格式转换
type ts转字符串 struct {
	D日期      *ts日期转字符串 // Date format "2006-01-02" | 日期格式
	Y年份      *ts日期转字符串 // Year just "2006" | 仅年份
	D月日      *ts日期转字符串 // Month-Day "01-02" | 月-日
	T时刻      *ts时间转字符串 // Time just "15:04:05" | 仅时间
	T时间      *ts时间转字符串 // Full datetime "2006-01-02 15:04:05" | 完整日期时间
	T日期和T和时刻 *ts时间转字符串 // ISO 8601 "2006-01-02T15:04:05" | ISO 8601 格式
	T时间14位数  *ts时间转字符串 // 14-digit "20060102150405" | 14位数字
	D日期8位数   *ts日期转字符串 // 8-digit date "20060102" | 8位数日期
	T时刻6位数   *ts时间转字符串 // 6-digit time "150405" | 6位数时间
}

// NewTs转字符串 creates a new time-to-string component with complete predefined formats
// Initializes various date/time format templates on common use cases
//
// NewTs转字符串 创建新的时间转字符串转换器，包含所有预定义格式
// 初始化各种常见用例的日期/时间格式模板
func NewTs转字符串() *ts转字符串 {
	return &ts转字符串{
		D日期:      NewTs日期转字符串("2006-01-02"),
		Y年份:      NewTs日期转字符串("2006"),
		D月日:      NewTs日期转字符串("01-02"),
		T时刻:      NewTs时间转字符串("15:04:05"),
		T时间:      NewTs时间转字符串("2006-01-02 15:04:05"),
		T日期和T和时刻: NewTs时间转字符串("2006-01-02T15:04:05"),
		T时间14位数:  NewTs时间转字符串("20060102150405"),
		D日期8位数:   NewTs日期转字符串("20060102"),
		T时刻6位数:   NewTs时间转字符串("150405"),
	}
}

// ts日期转字符串 formats dates using a specific layout pattern
// Holds the format layout and provides conversion methods
//
// ts日期转字符串 使用特定布局模式格式化日期
// 保存格式布局并提供转换方法
type ts日期转字符串 struct {
	s格式布局 string // Format layout pattern (e.g., "2006-01-02") | 格式布局模式
}

// NewTs日期转字符串 creates a new date component with specified layout
// s格式布局: Go time format layout (e.g., "2006-01-02")
//
// NewTs日期转字符串 创建新的日期格式化器，使用指定布局
// s格式布局: Go 时间格式布局（例如 "2006-01-02"）
func NewTs日期转字符串(s格式布局 string) *ts日期转字符串 {
	return &ts日期转字符串{
		s格式布局: s格式布局,
	}
}

// Get转字符串 converts a time object to string using the format layout
// t: the time object to convert
// Returns formatted date string
//
// Get转字符串 使用格式布局将时间对象转换为字符串
// t: 要转换的时间对象
// 返回格式化的日期字符串
func (ts *ts日期转字符串) Get转字符串(t time.Time) string {
	return t.Format(ts.s格式布局)
}

// Get现在日期 gets the current date string in specified timezone
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns current date as formatted string
//
// Get现在日期 获取指定时区的当前日期字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回格式化的当前日期字符串
func (ts *ts日期转字符串) Get现在日期(z时区位置 *time.Location) string {
	return ts.Get转字符串(time.Now().In(z时区位置))
}

// Get当日日期 gets the current date string in specified timezone
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns current date as formatted string
//
// Get当日日期 获取指定时区的今日日期字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回格式化的今日日期字符串
func (ts *ts日期转字符串) Get当日日期(z时区位置 *time.Location) string {
	return ts.Get现在日期(z时区位置)
}

// Get昨日日期 gets the-past-date's date string in specified timezone
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns the-past-date's date as formatted string
//
// Get昨日日期 获取指定时区的昨日日期字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回格式化的昨日日期字符串
func (ts *ts日期转字符串) Get昨日日期(z时区位置 *time.Location) string {
	return ts.Get转字符串(time.Now().In(z时区位置).AddDate(0, 0, -1))
}

// Get最近含当日的前N个工作日 gets recent N workdates including the-moment in specified timezone
// n: count of workdates to get (including the-moment if it is a work-date)
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns slice of date strings in reverse time-based sequence
//
// Get最近含当日的前N个工作日 获取指定时区包括今天在内的最近N个工作日
// n: 要获取的工作日数量（如果今天是工作日则包括今天）
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回按时间倒序排列的日期字符串切片
func (ts *ts日期转字符串) Get最近含当日的前N个工作日(n int, z时区位置 *time.Location) (workdates []string) {
	return ts.Get某天含当日的前N个工作日(time.Now().In(z时区位置), n)
}

// Get某天含当日的前N个工作日 gets recent N workdates from a specific date (including that date)
// t: the reference time to count backwards from
// n: count of workdates to get (including t if it is a work-date)
// Returns slice of date strings in reverse time-based sequence
//
// Get某天含当日的前N个工作日 从指定日期（包括该日期）向前获取N个工作日
// t: 开始回溯的参考时间
// n: 要获取的工作日数量（如果t是工作日则包括t）
// 返回按时间倒序排列的日期字符串切片
func (ts *ts日期转字符串) Get某天含当日的前N个工作日(t time.Time, n int) (workdates []string) {
	for {
		if WK.Is工作日(t) {
			workdates = append(workdates, ts.Get转字符串(t))
			if len(workdates) == n {
				return workdates
			}
		}
		t = t.AddDate(0, 0, -1)
	}
}

// Get获取当月1号的日期 gets the first date of the month from a given time
// t: the time object containing the month
// Returns formatted date string on the 1st date of that month
//
// Get获取当月1号的日期 从给定时间获取该月第一天的日期
// t: 包含月份信息的时间对象
// 返回该月1号的格式化日期字符串
func (ts *ts日期转字符串) Get获取当月1号的日期(t time.Time) string {
	Y, M := TI.Get年份和月份(t)
	return IS.Get日期(Y, M, 1)
}

// ts时间转字符串 formats time (datetime) using a specific layout pattern
// Holds the format layout and provides conversion methods
//
// ts时间转字符串 使用特定布局模式格式化时间（日期时间）
// 保存格式布局并提供转换方法
type ts时间转字符串 struct {
	s格式布局 string // Format layout pattern (e.g., "2006-01-02 15:04:05") | 格式布局模式
}

// NewTs时间转字符串 creates a new time (datetime) component with specified layout
// s格式布局: Go time format layout (e.g., "2006-01-02 15:04:05")
//
// NewTs时间转字符串 创建新的时间（日期时间）格式化器，使用指定布局
// s格式布局: Go 时间格式布局（例如 "2006-01-02 15:04:05"）
func NewTs时间转字符串(s格式布局 string) *ts时间转字符串 {
	return &ts时间转字符串{
		s格式布局: s格式布局,
	}
}

// Get转字符串 converts a time object to string using the format layout
// t: the time object to convert
// Returns formatted datetime string
//
// Get转字符串 使用格式布局将时间对象转换为字符串
// t: 要转换的时间对象
// 返回格式化的日期时间字符串
func (ts *ts时间转字符串) Get转字符串(t time.Time) string {
	return t.Format(ts.s格式布局)
}

// Get现在时间 gets the current datetime string in specified timezone
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns current datetime as formatted string
//
// Get现在时间 获取指定时区的当前日期时间字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回格式化的当前日期时间字符串
func (ts *ts时间转字符串) Get现在时间(z时区位置 *time.Location) string {
	return ts.Get转字符串(time.Now().In(z时区位置))
}
