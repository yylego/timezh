package timezh

import (
	"time"

	"github.com/yylego/erero"
)

// TS is the time-to-string component with predefined format templates
// Provides various date/time formatting options through structured fields
//
// TS 是全局时间转字符串转换器，包含预定义的格式模板
// 通过结构化字段提供各种日期/时间格式化选项
var TS = newTs转字符串()

// ts转字符串 holds various time-to-string components with different layouts
// Each field provides specific date/time format conversion
//
// ts转字符串 包含各种不同布局的时间转字符串格式化器
// 每个字段提供特定的日期/时间格式转换
type ts转字符串 struct {
	D日期      *ts日期转字符串 // Date format "2006-01-02" | 日期格式
	Y年份      *ts日期转字符串 // Just the 4 digits "2006" | 仅年份
	D月日      *ts日期转字符串 // Month and date "01-02" | 月-日
	T时刻      *ts时间转字符串 // Time just "15:04:05" | 仅时间
	T时分      *ts时间转字符串 // Clock time to the minute "15:04" | 时和分
	T时间      *ts时间转字符串 // Date and time "2006-01-02 15:04:05" | 完整日期时间
	T日期分时    *ts时间转字符串 // Datetime without seconds "2006-01-02 15:04" | 不带秒的日期时间
	T国际格式时间  *ts时间转字符串 // ISO 8601, the date and time split with the T: "2006-01-02T15:04:05" | ISO 8601 格式，日期和时刻用 T 隔开
	T十四位数字时间 *ts时间转字符串 // 14-digit "20060102150405" | 14位数字
	D八位数字日期  *ts日期转字符串 // 8-digit date "20060102" | 8位数日期
	T六位数字时刻  *ts时间转字符串 // 6-digit time "150405" | 6位数时间
}

// newTs转字符串 creates a new time-to-string component with complete predefined formats
// Initializes various date/time format templates on common use cases
//
// newTs转字符串 创建新的时间转字符串转换器，包含所有预定义格式
// 初始化各种常见用例的日期/时间格式模板
func newTs转字符串() *ts转字符串 {
	return &ts转字符串{
		D日期:      newTs日期转字符串(time.DateOnly),
		Y年份:      newTs日期转字符串("2006"),
		D月日:      newTs日期转字符串("01-02"),
		T时刻:      newTs时间转字符串(time.TimeOnly),
		T时分:      newTs时间转字符串("15:04"),
		T时间:      newTs时间转字符串(time.DateTime),
		T日期分时:    newTs时间转字符串("2006-01-02 15:04"),
		T国际格式时间:  newTs时间转字符串("2006-01-02T15:04:05"),
		T十四位数字时间: newTs时间转字符串("20060102150405"),
		D八位数字日期:  newTs日期转字符串("20060102"),
		T六位数字时刻:  newTs时间转字符串("150405"),
	}
}

// ts日期转字符串 formats, parses and checks dates in one layout
// The date arithmetic below needs a layout that holds the whole date, such as D日期 and D八位数字日期
//
// ts日期转字符串 按一种布局做日期的格式化、解析和校验
// 下面的日期加减、区间要用带年月日的布局，比如 D日期、D八位数字日期
type ts日期转字符串 struct {
	v格式 *rs时间格式
}

// newTs日期转字符串 creates a date component with the given layout
// s格式布局: Go time layout made of fixed-width numbers (e.g., "2006-01-02")
//
// newTs日期转字符串 创建日期格式化器，使用指定布局
// s格式布局: 位数固定的 Go 时间布局（例如 "2006-01-02"）
func newTs日期转字符串(s格式布局 string) *ts日期转字符串 {
	return &ts日期转字符串{
		v格式: newRs布局格式(s格式布局),
	}
}

// Get转为文本 converts a time object to string in the layout
//
// Get转为文本 按布局把时间对象转为字符串
func (ts *ts日期转字符串) Get转为文本(t time.Time) string {
	return t.Format(ts.v格式.s格式布局)
}

// Get转为时间 parses a string in the layout into a time in the zone, the shape must be the same as the layout
// In a zone whose clocks jump past 00:00 on that date, Go gives back an instant that can land on the date before;
// the TZ zones have fixed offsets, so this does not happen with them
//
// Get转为时间 按布局把字符串解析为该时区的时间，格式必须完全一致
// 某些时区在那一天会跳过 00:00（夏令时切换），Go 返回的时刻可能落在前一天；TZ 里的时区都是固定偏移，碰不到这种情况
func (ts *ts日期转字符串) Get转为时间(s string, z时区位置 *time.Location) (time.Time, error) {
	return ts.v格式.parse(s, z时区位置)
}

// Match格式 checks that the string is in the layout and the date exists
//
// Match格式 检查字符串是否符合布局、日期是否真实存在
func (ts *ts日期转字符串) Match格式(s string) bool {
	return ts.v格式.match(s)
}

// Get当日日期 gets the current date string in the zone
//
// Get当日日期 获取指定时区的今日日期字符串
func (ts *ts日期转字符串) Get当日日期(z时区位置 *time.Location) string {
	return ts.Get转为文本(time.Now().In(z时区位置))
}

// Get昨日日期 gets the date string before the current date in the zone
//
// Get昨日日期 获取指定时区的昨日日期字符串
func (ts *ts日期转字符串) Get昨日日期(z时区位置 *time.Location) string {
	return ts.Get转为文本(time.Now().In(z时区位置).AddDate(0, 0, -1))
}

// Get明日日期 gets the date string past the current date in the zone
//
// Get明日日期 获取指定时区的明日日期字符串
func (ts *ts日期转字符串) Get明日日期(z时区位置 *time.Location) string {
	return ts.Get转为文本(time.Now().In(z时区位置).AddDate(0, 0, 1))
}

// Get加减天数日期 moves a date string n days, back when n is negative, and gives it in the same layout
// D日期.Get加减天数日期("2025-03-01", -1) gives "2025-02-28"
//
// Get加减天数日期 把日期字符串往后挪 n 天，n 为负数时往前挪，结果用同一布局
// D日期.Get加减天数日期("2025-03-01", -1) 得到 "2025-02-28"
func (ts *ts日期转字符串) Get加减天数日期(s日期 string, n天数 int) (string, error) {
	v时间, err := ts.Get转为时间(s日期, time.UTC)
	if err != nil {
		return "", erero.WithMessage(err, "日期格式错误")
	}
	return ts.Get转为文本(v时间.AddDate(0, 0, n天数)), nil
}

// Get区间每天日期 lists each date from the start to the end, both included, oldest first
// It gives no dates when the start is past the end
//
// Get区间每天日期 列出从开始到结束的每一天，两头都包含，旧的在前
// 开始日期晚于结束日期时返回空
func (ts *ts日期转字符串) Get区间每天日期(s开始日期 string, e结束日期 string) ([]string, error) {
	t开始, err := ts.Get转为时间(s开始日期, time.UTC)
	if err != nil {
		return nil, erero.WithMessage(err, "开始日期格式错误")
	}
	t结束, err := ts.Get转为时间(e结束日期, time.UTC)
	if err != nil {
		return nil, erero.WithMessage(err, "结束日期格式错误")
	}
	var dates []string
	for t := t开始; !t.After(t结束); t = t.AddDate(0, 0, 1) {
		dates = append(dates, ts.Get转为文本(t))
	}
	return dates, nil
}

// Get截至当日工作日期 gets the latest n Mon-Fri dates up to now in the zone, newest first
//
// Get截至当日工作日期 获取指定时区包括今天在内的最近 n 个工作日（周一到周五），新的在前
func (ts *ts日期转字符串) Get截至当日工作日期(n int, z时区位置 *time.Location) []string {
	return ts.Get截至某日工作日期(time.Now().In(z时区位置), n)
}

// Get截至某日工作日期 gets the latest n Mon-Fri dates up to t, newest first
//
// Get截至某日工作日期 从 t 那天（含）往前取 n 个工作日（周一到周五），新的在前，n 不大于 0 时返回空
func (ts *ts日期转字符串) Get截至某日工作日期(t time.Time, n int) []string {
	var dates []string
	for len(dates) < n {
		if WK.Is工作日(t) {
			dates = append(dates, ts.Get转为文本(t))
		}
		t = t.AddDate(0, 0, -1)
	}
	return dates
}

// Get月初日期 gets the 1st of the month of t, such as "2025-05-01" with D日期 and "20250501" with D八位数字日期
//
// Get月初日期 获取 t 所在月份的 1 号，比如 D日期 得到 "2025-05-01"，D八位数字日期 得到 "20250501"
func (ts *ts日期转字符串) Get月初日期(t time.Time) string {
	return ts.Get某月月初日期(t, 0)
}

// Get月末日期 gets the last date of the month of t, such as "2024-02-29"
//
// Get月末日期 获取 t 所在月份的最后一天，比如 "2024-02-29"
func (ts *ts日期转字符串) Get月末日期(t time.Time) string {
	return ts.Get某月月末日期(t, 0)
}

// Get某月月初日期 gets the 1st of the month n months from the month of t, -1 gives the month before
// It counts on months, so 03-31 with -1 gives 02-01, not the 03-03 that t.AddDate(0, -1, 0) lands on
//
// Get某月月初日期 获取 t 所在月份往后第 n 个月的 1 号，n 为 -1 就是上个月
// 按月份计算，所以 03-31 传 -1 得到 02-01，不会像 t.AddDate(0, -1, 0) 那样落到 03-03
func (ts *ts日期转字符串) Get某月月初日期(t time.Time, n月数 int) string {
	return ts.Get转为文本(time.Date(t.Year(), t.Month()+time.Month(n月数), 1, 0, 0, 0, 0, t.Location()))
}

// Get某月月末日期 gets the last date of the month n months from the month of t, -1 gives the month before
//
// Get某月月末日期 获取 t 所在月份往后第 n 个月的最后一天，n 为 -1 就是上个月
func (ts *ts日期转字符串) Get某月月末日期(t time.Time, n月数 int) string {
	return ts.Get转为文本(time.Date(t.Year(), t.Month()+time.Month(n月数)+1, 0, 0, 0, 0, 0, t.Location()))
}

// Get当周周一日期 gets the Mon of the week of t, the week runs from Mon to Sun
//
// Get当周周一日期 获取 t 所在那一周的周一，一周从周一算到周日
func (ts *ts日期转字符串) Get当周周一日期(t time.Time) string {
	return ts.Get转为文本(t.AddDate(0, 0, -daysFromMon(t)))
}

// Get当周周日日期 gets the Sun of the week of t, the week runs from Mon to Sun
//
// Get当周周日日期 获取 t 所在那一周的周日，一周从周一算到周日
func (ts *ts日期转字符串) Get当周周日日期(t time.Time) string {
	return ts.Get转为文本(t.AddDate(0, 0, 6-daysFromMon(t)))
}

// daysFromMon gives the count of days from the Mon of the week to t: 0 on Mon, 6 on Sun
// daysFromMon 返回 t 比当周周一晚几天：周一是 0，周日是 6
func daysFromMon(t time.Time) int {
	return (int(t.Weekday()) + 6) % 7
}

// ts时间转字符串 formats, parses and checks times in one layout
//
// ts时间转字符串 按一种布局做时间的格式化、解析和校验
type ts时间转字符串 struct {
	v格式 *rs时间格式
}

// newTs时间转字符串 creates a time component with the given layout
// s格式布局: Go time layout made of fixed-width numbers (e.g., "2006-01-02 15:04:05")
//
// newTs时间转字符串 创建时间格式化器，使用指定布局
// s格式布局: 位数固定的 Go 时间布局（例如 "2006-01-02 15:04:05"）
func newTs时间转字符串(s格式布局 string) *ts时间转字符串 {
	return &ts时间转字符串{
		v格式: newRs布局格式(s格式布局),
	}
}

// Get转为文本 converts a time object to string in the layout
//
// Get转为文本 按布局把时间对象转为字符串
func (ts *ts时间转字符串) Get转为文本(t time.Time) string {
	return t.Format(ts.v格式.s格式布局)
}

// Get转为时间 parses a string in the layout into a time in the zone, the shape must be the same as the layout
//
// Get转为时间 按布局把字符串解析为该时区的时间，格式必须完全一致
func (ts *ts时间转字符串) Get转为时间(s string, z时区位置 *time.Location) (time.Time, error) {
	return ts.v格式.parse(s, z时区位置)
}

// Match格式 checks that the string is in the layout and the time exists
//
// Match格式 检查字符串是否符合布局、时间是否真实存在
func (ts *ts时间转字符串) Match格式(s string) bool {
	return ts.v格式.match(s)
}

// Get现在时间 gets the current time string in the zone
//
// Get现在时间 获取指定时区的当前时间字符串
func (ts *ts时间转字符串) Get现在时间(z时区位置 *time.Location) string {
	return ts.Get转为文本(time.Now().In(z时区位置))
}
