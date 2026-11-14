package timezh

import (
	"time"

	"github.com/yylego/erero"
)

// NT is the numeric-to-time conversion component
// Provides methods to create time objects from numeric components
//
// NT 是全局数字转时间转换器
// 提供从数字组件创建时间对象的方法
var NT = &nt把数字转化为时间{}

// nt把数字转化为时间 constructs time objects from numeric date/time components
// Numbers that do not make an existing date/time give an error, unlike time.Date, which moves them on to the next month and date
//
// nt把数字转化为时间 从数字日期/时间组件构建时间对象
// 拼不成真实日期/时刻的数字返回错误，不像 time.Date 那样顺延到下个月、下一天
type nt把数字转化为时间 struct{}

// New当日时刻时间 creates a time object at the current date with specified time values
// v几时: the time-value (0-23)
// v几分: minute (0-59)
// v几秒: second (0-59)
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns time object at the current date with specified time
//
// New当日时刻时间 创建今天日期、指定时分秒的时间对象
// v几时: 小时（0-23）
// v几分: 分钟（0-59）
// v几秒: 秒数（0-59）
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回今天日期、指定时分秒的时间对象
func (nt *nt把数字转化为时间) New当日时刻时间(v几时, v几分, v几秒 int, z时区位置 *time.Location) (time.Time, error) {
	if _, err := IS.Get时刻(v几时, v几分, v几秒); err != nil {
		return time.Time{}, erero.WithMessage(err, "时分秒不对")
	}
	now := time.Now().In(z时区位置)
	return time.Date(now.Year(), now.Month(), now.Day(), v几时, v几分, v几秒, 0, z时区位置), nil
}

// New当日零点时间 gives 00:00:00 of the date of t, in the zone of t
//
// New当日零点时间 返回 t 那一天的 00:00:00，时区跟 t 一样
func (nt *nt把数字转化为时间) New当日零点时间(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// New某日零点时间 creates a time object at midnight on specified date
// v年份: the date-element (e.g., 2025)
// v月份: month (1-12)
// v几号: date of month (1-31)
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
// Returns time object at 00:00:00 on the specified date
//
// New某日零点时间 创建指定日期零时（午夜）的时间对象
// v年份: 年份（例如 2025）
// v月份: 月份（1-12）
// v几号: 日期（1-31）
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
// 返回指定日期 00:00:00 的时间对象
func (nt *nt把数字转化为时间) New某日零点时间(v年份, v月份, v几号 int, z时区位置 *time.Location) (time.Time, error) {
	if _, err := IS.Get日期(v年份, v月份, v几号); err != nil {
		return time.Time{}, erero.WithMessage(err, "年月日不对")
	}
	return time.Date(v年份, time.Month(v月份), v几号, 0, 0, 0, 0, z时区位置), nil
}
