package timezh

import "time"

// NS is the timestamp-to-time conversion component
// Provides methods to convert Unix timestamps into time objects
//
// NS 是全局时间戳转时间转换器
// 提供将 Unix 时间戳转换为时间对象的方法
var NS = &ns把时间戳转为时间{}

// ns把时间戳转为时间 converts Unix timestamps to time objects
// Supports nanosecond, microsecond, millisecond, and second timestamp formats
//
// ns把时间戳转为时间 将 Unix 时间戳转换为时间对象
// 支持纳秒、微秒、毫秒和秒级时间戳格式
type ns把时间戳转为时间 struct{}

// Get纳秒时间 converts nanosecond timestamp to time object
// ns纳秒数: Unix timestamp in nanoseconds
// Returns corresponding time object
//
// Get纳秒时间 将纳秒级时间戳转换为时间对象
// ns纳秒数: 纳秒级 Unix 时间戳
// 返回对应的时间对象
func (T *ns把时间戳转为时间) Get纳秒时间(ns纳秒数 int64) time.Time {
	return time.Unix(ns纳秒数/1e9, ns纳秒数%1e9)
}

// Get微秒时间 converts microsecond timestamp to time object
// ms微秒数: Unix timestamp in microseconds
// Returns corresponding time object
//
// Get微秒时间 将微秒级时间戳转换为时间对象
// ms微秒数: 微秒级 Unix 时间戳
// 返回对应的时间对象
func (T *ns把时间戳转为时间) Get微秒时间(ms微秒数 int64) time.Time {
	return time.UnixMicro(ms微秒数)
}

// Get毫秒时间 converts millisecond timestamp to time object
// ms毫秒数: Unix timestamp in milliseconds
// Returns corresponding time object
//
// Get毫秒时间 将毫秒级时间戳转换为时间对象
// ms毫秒数: 毫秒级 Unix 时间戳
// 返回对应的时间对象
func (T *ns把时间戳转为时间) Get毫秒时间(ms毫秒数 int64) time.Time {
	return time.UnixMilli(ms毫秒数)
}

// Get秒数时间 converts second timestamp to time object
// sn总秒数: Unix timestamp in seconds
// Returns corresponding time object
//
// Get秒数时间 将秒级时间戳转换为时间对象
// sn总秒数: 秒级 Unix 时间戳
// 返回对应的时间对象
func (T *ns把时间戳转为时间) Get秒数时间(sn总秒数 int64) time.Time {
	return time.Unix(sn总秒数, 0)
}
