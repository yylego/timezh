package timezh

import (
	"fmt"
	"strconv"

	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// SI is the string-to-numeric component
// Provides methods to parse time strings into numeric components
//
// SI 是全局字符串转整数转换器
// 提供将时间字符串解析为数字组件的方法
var SI = &si把字符串时间转换整数{}

// si把字符串时间转换整数 parses time strings into numeric components
// Extracts date-element, month, date, time-value, minute, second values from time strings
//
// si把字符串时间转换整数 将时间字符串解析为数字组件
// 从时间字符串提取年、月、日、时、分、秒数值
type si把字符串时间转换整数 struct{}

// Get通过日期得到年份 extracts date-element from a date string
// s标准日期: date string in format "2006-01-02"
// Returns date-element value, and/if 0 if parsing fails
//
// Get通过日期得到年份 从日期字符串提取年份
// s标准日期: "2006-01-02" 格式的日期字符串
// 返回年份值，解析失败时返回 0
func (si *si把字符串时间转换整数) Get通过日期得到年份(s标准日期 string) (v年份 int) {
	if !RS.Match日期(s标准日期) {
		zaplog.LOG.Warn("日期格式不对", zap.String("s", s标准日期))
		return 0
	}
	v年份, err := strconv.Atoi(s标准日期[0:4])
	if err != nil {
		zaplog.LOG.Warn("日期格式不对", zap.String("s", s标准日期), zap.Error(err))
		return 0
	}
	return v年份
}

// Get通过日期得到年份月份哪日 extracts date-element, month, and date from a date string
// s标准日期: date string in format "2006-01-02"
// Returns date-element, month (1-12), date (1-31), and/if zeros if parsing fails
//
// Get通过日期得到年份月份哪日 从日期字符串提取年、月、日
// s标准日期: "2006-01-02" 格式的日期字符串
// 返回年份、月份（1-12）、日期（1-31），解析失败时返回零值
func (si *si把字符串时间转换整数) Get通过日期得到年份月份哪日(s标准日期 string) (v年份, v月份, v几号 int) {
	num, err := fmt.Sscanf(s标准日期, "%4d-%2d-%2d", &v年份, &v月份, &v几号)
	if err != nil {
		zaplog.LOG.Warn("日期格式不对", zap.String("s", s标准日期), zap.Error(err))
		return 0, 0, 0
	}
	if num != 3 {
		zaplog.LOG.Warn("日期格式不对", zap.String("s", s标准日期), zap.Int("num", num))
		return 0, 0, 0
	}
	return v年份, v月份, v几号
}

// Get通过时刻得到几时几分秒数 extracts time-value, minute, and second from a time-of-moment string
// s标准时刻: time-of-moment string in format "15:04:05"
// Returns time-value (0-23), minute (0-59), second (0-59), and/if zeros if parsing fails
//
// Get通过时刻得到几时几分秒数 从时刻字符串提取时、分、秒
// s标准时刻: "15:04:05" 格式的时刻字符串
// 返回小时（0-23）、分钟（0-59）、秒数（0-59），解析失败时返回零值
func (si *si把字符串时间转换整数) Get通过时刻得到几时几分秒数(s标准时刻 string) (v几时, v几分, v秒数 int) {
	num, err := fmt.Sscanf(s标准时刻, "%2d:%2d:%2d", &v几时, &v几分, &v秒数)
	if err != nil {
		zaplog.LOG.Warn("时刻格式不对", zap.String("s", s标准时刻), zap.Error(err))
		return 0, 0, 0
	}
	if num != 3 {
		zaplog.LOG.Warn("时刻格式不对", zap.String("s", s标准时刻), zap.Int("num", num))
		return 0, 0, 0
	}
	return v几时, v几分, v秒数
}

// Get通过时间得到整数 extracts complete date/time components from a datetime string
// s标准时间: datetime string in format "2006-01-02 15:04:05"
// Returns date-element, month (1-12), date (1-31), time-value (0-23), minute (0-59), second (0-59), and/if zeros if parsing fails
//
// Get通过时间得到整数 从日期时间字符串提取所有日期/时间组件
// s标准时间: "2006-01-02 15:04:05" 格式的日期时间字符串
// 返回年份、月份（1-12）、日期（1-31）、小时（0-23）、分钟（0-59）、秒数（0-59），解析失败时返回零值
func (si *si把字符串时间转换整数) Get通过时间得到整数(s标准时间 string) (v年份, v月份, v几号, v几时, v几分, v秒数 int) {
	if !RS.Match时间(s标准时间) {
		zaplog.LOG.Warn("时间格式不对", zap.String("s", s标准时间))
		return 0, 0, 0, 0, 0, 0
	}
	num, err := fmt.Sscanf(s标准时间, "%4d-%2d-%2d %2d:%2d:%2d", &v年份, &v月份, &v几号, &v几时, &v几分, &v秒数)
	if err != nil {
		zaplog.LOG.Warn("时间格式不对", zap.String("s", s标准时间), zap.Error(err))
		return 0, 0, 0, 0, 0, 0
	}
	if num != 6 {
		zaplog.LOG.Warn("时间格式不对", zap.String("s", s标准时间), zap.Int("num", num))
		return 0, 0, 0, 0, 0, 0
	}
	return v年份, v月份, v几号, v几时, v几分, v秒数
}
