package timezh

import (
	"time"

	"github.com/yylego/erero"
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

// Get年份 extracts the 4 leading digits of the date as an int, such as 2021
// s标准日期: date string in format "2006-01-02"
//
// Get年份 从日期字符串提取年份
// s标准日期: "2006-01-02" 格式的日期字符串
func (si *si把字符串时间转换整数) Get年份(s标准日期 string) (v年份 int, err error) {
	v时间, err := TS.D日期.Get转为时间(s标准日期, time.UTC)
	if err != nil {
		return 0, erero.WithMessage(err, "日期格式不对")
	}
	return v时间.Year(), nil
}

// Get拆分日期 extracts the three numbers of a date string
// s标准日期: date string in format "2006-01-02"
//
// Get拆分日期 从日期字符串提取年、月、日
// s标准日期: "2006-01-02" 格式的日期字符串
func (si *si把字符串时间转换整数) Get拆分日期(s标准日期 string) (v年份, v月份, v几号 int, err error) {
	v时间, err := TS.D日期.Get转为时间(s标准日期, time.UTC)
	if err != nil {
		return 0, 0, 0, erero.WithMessage(err, "日期格式不对")
	}
	return v时间.Year(), int(v时间.Month()), v时间.Day(), nil
}

// Get拆分时刻 extracts the three numbers of a time-of-moment string
// s标准时刻: time-of-moment string in format "15:04:05"
//
// Get拆分时刻 从时刻字符串提取时、分、秒
// s标准时刻: "15:04:05" 格式的时刻字符串
func (si *si把字符串时间转换整数) Get拆分时刻(s标准时刻 string) (v几时, v几分, v秒数 int, err error) {
	v时间, err := TS.T时刻.Get转为时间(s标准时刻, time.UTC)
	if err != nil {
		return 0, 0, 0, erero.WithMessage(err, "时刻格式不对")
	}
	return v时间.Hour(), v时间.Minute(), v时间.Second(), nil
}

// Get拆分时间 extracts each date/time component from a datetime string
// s标准时间: datetime string in format "2006-01-02 15:04:05"
//
// Get拆分时间 从日期时间字符串提取所有日期/时间组件
// s标准时间: "2006-01-02 15:04:05" 格式的日期时间字符串
func (si *si把字符串时间转换整数) Get拆分时间(s标准时间 string) (v年份, v月份, v几号, v几时, v几分, v秒数 int, err error) {
	v时间, err := TS.T时间.Get转为时间(s标准时间, time.UTC)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, erero.WithMessage(err, "时间格式不对")
	}
	return v时间.Year(), int(v时间.Month()), v时间.Day(), v时间.Hour(), v时间.Minute(), v时间.Second(), nil
}
