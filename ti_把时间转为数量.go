package timezh

import "time"

// TI is the time-to-numeric component
// Provides methods to extract numeric components from time objects
//
// TI 是全局时间转数量转换器
// 提供从时间对象提取数字组件的方法
var TI = &ti把时间转为数量{}

// ti把时间转为数量 extracts numeric components from time objects
// Converts time objects into date-element, month, date, time-value, minute, second values
//
// ti把时间转为数量 从时间对象提取数字组件
// 将时间对象转换为年、月、日、时、分、秒数值
type ti把时间转为数量 struct{}

// Get年份月份几号 extracts date-element, month, and date from time object
// t: the time object to extract from
// Returns date-element, month (1-12), and date (1-31)
//
// Get年份月份几号 从时间对象提取年、月、日
// t: 要提取的时间对象
// 返回年份、月份（1-12）、日期（1-31）
func (T *ti把时间转为数量) Get年份月份几号(t time.Time) (v年份, v月份, v几号 int) {
	v年份 = t.Year()
	v月份 = int(t.Month())
	v几号 = t.Day()
	return
}

// Get几时几分秒数 extracts time-value, minute, and second from time object
// t: the time object to extract from
// Returns time-value (0-23), minute (0-59), and second (0-59)
//
// Get几时几分秒数 从时间对象提取时、分、秒
// t: 要提取的时间对象
// 返回小时（0-23）、分钟（0-59）、秒数（0-59）
func (T *ti把时间转为数量) Get几时几分秒数(t time.Time) (v几时, v几分, v秒数 int) {
	v几时 = t.Hour()
	v几分 = t.Minute()
	v秒数 = t.Second()
	return
}

// Get数量 extracts complete date/time components from time object
// t: the time object to extract from
// Returns date-element, month (1-12), date (1-31), time-value (0-23), minute (0-59), second (0-59)
//
// Get数量 从时间对象提取所有日期/时间组件
// t: 要提取的时间对象
// 返回年份、月份（1-12）、日期（1-31）、小时（0-23）、分钟（0-59）、秒数（0-59）
func (T *ti把时间转为数量) Get数量(t time.Time) (v年份, v月份, v几号, v几时, v几分, v秒数 int) {
	v年份 = t.Year()
	v月份 = int(t.Month())
	v几号 = t.Day()
	v几时 = t.Hour()
	v几分 = t.Minute()
	v秒数 = t.Second()
	return
}

// Get年份和月份 extracts date-element and month from time object
// t: the time object to extract from
// Returns date-element and month (1-12)
//
// Get年份和月份 从时间对象提取年份和月份
// t: 要提取的时间对象
// 返回年份和月份（1-12）
func (T *ti把时间转为数量) Get年份和月份(t time.Time) (v年份, v月份 int) {
	v年份 = t.Year()
	v月份 = int(t.Month())
	return
}

// Get年份 extracts date-element from time object
// t: the time object to extract from
// Returns date-element
//
// Get年份 从时间对象提取年份
// t: 要提取的时间对象
// 返回年份
func (T *ti把时间转为数量) Get年份(t time.Time) (v年份 int) {
	v年份 = t.Year()
	return
}
