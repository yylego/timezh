package timezh

import (
	"time"

	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// ST is the the string-to-time component
// Provides methods to parse various time string formats into time objects
//
// ST 是全局字符串转时间转换器
// 提供将各种时间字符串格式解析为时间对象的方法
var ST = &st把字符串转为时间{}

// st把字符串转为时间 parses time strings into time objects
// Supports various date/time string formats with timezone handling
//
// st把字符串转为时间 将时间字符串解析为时间对象
// 支持各种日期/时间字符串格式，带时区处理
type st把字符串转为时间 struct{}

// Get时间 parses a standard datetime string into time object
// s标准时间: datetime string in format "2006-01-02 15:04:05"
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns parsed time object, and/if Unix epoch time if parsing fails
//
// Get时间 将标准日期时间字符串解析为时间对象
// s标准时间: "2006-01-02 15:04:05" 格式的日期时间字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回解析的时间对象，解析失败时返回 Unix 纪元时间
func (st *st把字符串转为时间) Get时间(s标准时间 string, z时区位置 *time.Location) time.Time {
	v时间, err := time.ParseInLocation("2006-01-02 15:04:05", s标准时间, z时区位置)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间), zap.Error(err))
		v时间 = time.Unix(0, 0)
	}
	return v时间
}

// Get日期和分时 parses a datetime string without seconds into time object
// s标准时间没有秒数: datetime string in format "2006-01-02 15:04"
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns parsed time object, and/if Unix epoch time if parsing fails
//
// Get日期和分时 将不含秒数的日期时间字符串解析为时间对象
// s标准时间没有秒数: "2006-01-02 15:04" 格式的日期时间字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回解析的时间对象，解析失败时返回 Unix 纪元时间
func (st *st把字符串转为时间) Get日期和分时(s标准时间没有秒数 string, z时区位置 *time.Location) time.Time {
	v时间, err := time.ParseInLocation("2006-01-02 15:04", s标准时间没有秒数, z时区位置)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间没有秒数), zap.Error(err))
		v时间 = time.Unix(0, 0)
	}
	return v时间
}

// Get通过日期得到时间 parses a date string into time object
// s标准日期: date string in format "2006-01-02"
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns parsed time object at midnight, and/if Unix epoch time if parsing fails
//
// Get通过日期得到时间 将日期字符串解析为时间对象
// s标准日期: "2006-01-02" 格式的日期字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回零时（午夜）的时间对象，解析失败时返回 Unix 纪元时间
func (st *st把字符串转为时间) Get通过日期得到时间(s标准日期 string, z时区位置 *time.Location) time.Time {
	v时间, err := time.ParseInLocation("2006-01-02", s标准日期, z时区位置)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准日期), zap.Error(err))
		v时间 = time.Unix(0, 0)
	}
	return v时间
}

// Get通过时刻得到当日时间 parses a time-of-moment string and creates the-moment's time object
// s标准时刻: time string in format "15:04:05"
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns the-moment's time at specified time-value/minute/second, and/if Unix epoch time if parsing fails
//
// Get通过时刻得到当日时间 解析时刻字符串并创建今天的时间对象
// s标准时刻: "15:04:05" 格式的时间字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回今天指定时分秒的时间对象，解析失败时返回 Unix 纪元时间
func (st *st把字符串转为时间) Get通过时刻得到当日时间(s标准时刻 string, z时区位置 *time.Location) time.Time {
	v时间, err := time.ParseInLocation("15:04:05", s标准时刻, z时区位置)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时刻), zap.Error(err))
		v时间 = time.Unix(0, 0)
	}
	return NT.New由当日时分秒得到时间(v时间.Hour(), v时间.Minute(), v时间.Second(), z时区位置)
}
