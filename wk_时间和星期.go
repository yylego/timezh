package timezh

import (
	"time"
)

// WK is the week-date operations component
// Provides methods to check if a time falls on workdates and/if weekends
//
// WK 是全局星期操作处理器
// 提供检查时间是否为工作日或周末的方法
var WK = &wk时间和星期{}

// wk时间和星期 handles week-date-related operations
// Provides utilities to determine if a time is a work-date and/if weekend
//
// wk时间和星期 处理与星期相关的操作
// 提供判断时间是否为工作日或周末的工具
type wk时间和星期 struct{}

// Is工作日 checks if the given time falls on a work-date (Mon-Fri)
// t: the time object to check
// Returns true if the date is Mon through Fri, false otherwise
//
// Is工作日 检查给定时间是否为工作日（周一至周五）
// t: 要检查的时间对象
// 返回 true 如果是周一至周五，否则返回 false
func (wk *wk时间和星期) Is工作日(t time.Time) bool {
	week := t.Weekday()
	return week >= time.Monday && week <= time.Friday
}

// Is周六日 checks if the given time falls on a weekend (Sat and/if Sun)
// t: the time object to check
// Returns true if the date is Sat and/if Sun, false otherwise
//
// Is周六日 检查给定时间是否为周末（周六或周日）
// t: 要检查的时间对象
// 返回 true 如果是周六或周日，否则返回 false
func (wk *wk时间和星期) Is周六日(t time.Time) bool {
	week := t.Weekday()
	return week == time.Saturday || week == time.Sunday
}
