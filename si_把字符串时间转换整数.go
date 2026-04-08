package timezh

import (
	"fmt"
	"strconv"

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

// Get通过日期得到年份 extracts date-element from a date string
// s标准日期: date string in format "2006-01-02"
//
// Get通过日期得到年份 从日期字符串提取年份
// s标准日期: "2006-01-02" 格式的日期字符串
func (si *si把字符串时间转换整数) Get通过日期得到年份(s标准日期 string) (v年份 int, err error) {
	if !RS.Match日期(s标准日期) {
		return 0, erero.Errorf("日期格式不对: %s", s标准日期)
	}
	v年份, err = strconv.Atoi(s标准日期[0:4])
	if err != nil {
		return 0, erero.WithMessage(err, "日期格式不对")
	}
	return v年份, nil
}

// Get通过日期得到年份月份哪日 extracts date-element, month, and date from a date string
// s标准日期: date string in format "2006-01-02"
//
// Get通过日期得到年份月份哪日 从日期字符串提取年、月、日
// s标准日期: "2006-01-02" 格式的日期字符串
func (si *si把字符串时间转换整数) Get通过日期得到年份月份哪日(s标准日期 string) (v年份, v月份, v几号 int, err error) {
	num, err := fmt.Sscanf(s标准日期, "%4d-%2d-%2d", &v年份, &v月份, &v几号)
	if err != nil {
		return 0, 0, 0, erero.WithMessage(err, "日期格式不对")
	}
	if num != 3 {
		return 0, 0, 0, erero.Errorf("日期格式不对: 解析到 %d 个字段, 期望 3 个", num)
	}
	return v年份, v月份, v几号, nil
}

// Get通过时刻得到几时几分秒数 extracts time-value, minute, and second from a time-of-moment string
// s标准时刻: time-of-moment string in format "15:04:05"
//
// Get通过时刻得到几时几分秒数 从时刻字符串提取时、分、秒
// s标准时刻: "15:04:05" 格式的时刻字符串
func (si *si把字符串时间转换整数) Get通过时刻得到几时几分秒数(s标准时刻 string) (v几时, v几分, v秒数 int, err error) {
	num, err := fmt.Sscanf(s标准时刻, "%2d:%2d:%2d", &v几时, &v几分, &v秒数)
	if err != nil {
		return 0, 0, 0, erero.WithMessage(err, "时刻格式不对")
	}
	if num != 3 {
		return 0, 0, 0, erero.Errorf("时刻格式不对: 解析到 %d 个字段, 期望 3 个", num)
	}
	return v几时, v几分, v秒数, nil
}

// Get通过时间得到整数 extracts complete date/time components from a datetime string
// s标准时间: datetime string in format "2006-01-02 15:04:05"
//
// Get通过时间得到整数 从日期时间字符串提取所有日期/时间组件
// s标准时间: "2006-01-02 15:04:05" 格式的日期时间字符串
func (si *si把字符串时间转换整数) Get通过时间得到整数(s标准时间 string) (v年份, v月份, v几号, v几时, v几分, v秒数 int, err error) {
	if !RS.Match时间(s标准时间) {
		return 0, 0, 0, 0, 0, 0, erero.Errorf("时间格式不对: %s", s标准时间)
	}
	num, err := fmt.Sscanf(s标准时间, "%4d-%2d-%2d %2d:%2d:%2d", &v年份, &v月份, &v几号, &v几时, &v几分, &v秒数)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, erero.WithMessage(err, "时间格式不对")
	}
	if num != 6 {
		return 0, 0, 0, 0, 0, 0, erero.Errorf("时间格式不对: 解析到 %d 个字段, 期望 6 个", num)
	}
	return v年份, v月份, v几号, v几时, v几分, v秒数, nil
}
