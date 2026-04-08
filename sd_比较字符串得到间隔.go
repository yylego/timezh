package timezh

import (
	"time"

	"github.com/yylego/erero"
)

// SD is the string duration component
// Provides methods to calculate time intervals between time strings
//
// SD 是全局字符串间隔计算器
// 提供计算时间字符串间时间间隔的方法
var SD = &sd比较字符串得到间隔{}

// sd比较字符串得到间隔 calculates time durations between time strings
// Computes intervals from now to dates, times, and between two time strings
//
// sd比较字符串得到间隔 计算时间字符串间的时间间隔
// 计算从现在到日期、时间，或两个时间字符串之间的间隔
type sd比较字符串得到间隔 struct{}

// Get现在到日期的间隔 calculates duration from now to a target date
// s标准日期: target date string in format "2006-01-02"
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
//
// Get现在到日期的间隔 计算从现在到目标日期的间隔
// s标准日期: "2006-01-02" 格式的目标日期字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
func (SD *sd比较字符串得到间隔) Get现在到日期的间隔(s标准日期 string, z时区位置 *time.Location) (time.Duration, error) {
	return SD.Get时间到日期的间隔(time.Now().In(z时区位置), s标准日期)
}

// Get时间到日期的间隔 calculates duration from a specific time to a target date
// t: the reference time
// s标准日期: target date string in format "2006-01-02"
//
// Get时间到日期的间隔 计算从指定时间到目标日期的间隔
// t: 参考时间
// s标准日期: "2006-01-02" 格式的目标日期字符串
func (SD *sd比较字符串得到间隔) Get时间到日期的间隔(t time.Time, s标准日期 string) (time.Duration, error) {
	v时间, err := ST.Get通过日期得到时间(s标准日期, t.Location())
	if err != nil {
		return 0, erero.WithMessage(err, "计算间隔失败")
	}
	return v时间.Sub(t), nil
}

// Get现在到当天时刻的间隔 calculates duration from now to a time-of-moment the-moment
// s标准时刻: target time-of-moment string in format "15:04:05"
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
//
// Get现在到当天时刻的间隔 计算从现在到今天某时刻的间隔
// s标准时刻: "15:04:05" 格式的目标时刻字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
func (SD *sd比较字符串得到间隔) Get现在到当天时刻的间隔(s标准时刻 string, z时区位置 *time.Location) (time.Duration, error) {
	return SD.Get时间到同天时刻的间隔(time.Now().In(z时区位置), s标准时刻)
}

// Get时间到同天时刻的间隔 calculates duration from a time to a time-of-moment on the same date
// t: the reference time
// s标准时刻: target time-of-moment string in format "15:04:05"
//
// Get时间到同天时刻的间隔 计算从某时间到同一天某时刻的间隔
// t: 参考时间
// s标准时刻: "15:04:05" 格式的目标时刻字符串
func (SD *sd比较字符串得到间隔) Get时间到同天时刻的间隔(t time.Time, s标准时刻 string) (time.Duration, error) {
	v几时, v几分, v秒数, err := SI.Get通过时刻得到几时几分秒数(s标准时刻)
	if err != nil {
		return 0, erero.WithMessage(err, "时刻格式错误")
	}
	v目标 := time.Date(t.Year(), t.Month(), t.Day(), v几时, v几分, v秒数, 0, t.Location())
	return v目标.Sub(t), nil
}

// Get现在到时间的间隔 calculates duration from now to a target datetime
// s标准时间: target datetime string in format "2006-01-02 15:04:05"
// z时区位置: the timezone location (e.g., TZ.CST, TZ.UTC)
//
// Get现在到时间的间隔 计算从现在到目标日期时间的间隔
// s标准时间: "2006-01-02 15:04:05" 格式的目标日期时间字符串
// z时区位置: 时区位置（例如 TZ.CST、TZ.UTC）
func (SD *sd比较字符串得到间隔) Get现在到时间的间隔(s标准时间 string, z时区位置 *time.Location) (time.Duration, error) {
	return SD.Get时间到时间的间隔(time.Now().In(z时区位置), s标准时间)
}

// Get时间到时间的间隔 calculates duration from a time to a target datetime
// t: the reference time
// s标准时间: target datetime string in format "2006-01-02 15:04:05"
//
// Get时间到时间的间隔 计算从某时间到目标日期时间的间隔
// t: 参考时间
// s标准时间: "2006-01-02 15:04:05" 格式的目标日期时间字符串
func (SD *sd比较字符串得到间隔) Get时间到时间的间隔(t time.Time, s标准时间 string) (time.Duration, error) {
	v时间, err := ST.Get时间(s标准时间, t.Location())
	if err != nil {
		return 0, erero.WithMessage(err, "计算间隔失败")
	}
	return v时间.Sub(t), nil
}

// Get两字符串时间间隔 calculates duration between two datetime strings
// s开始时间: start datetime string in format "2006-01-02 15:04:05"
// e结束时间: end datetime string in format "2006-01-02 15:04:05"
//
// Get两字符串时间间隔 计算两个日期时间字符串之间的间隔
// s开始时间: "2006-01-02 15:04:05" 格式的开始日期时间字符串
// e结束时间: "2006-01-02 15:04:05" 格式的结束日期时间字符串
func (SD *sd比较字符串得到间隔) Get两字符串时间间隔(s开始时间, e结束时间 string) (time.Duration, error) {
	t开始, err := ST.Get时间(s开始时间, time.UTC)
	if err != nil {
		return 0, erero.WithMessage(err, "开始时间格式错误")
	}
	t结束, err := ST.Get时间(e结束时间, time.UTC)
	if err != nil {
		return 0, erero.WithMessage(err, "结束时间格式错误")
	}
	return t结束.Sub(t开始), nil
}
