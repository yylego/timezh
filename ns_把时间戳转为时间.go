package timezh

import "time"

// NS is the timestamp-to-time conversion component
// Provides methods to convert Unix timestamps into time objects
//
// NS 是全局时间戳转时间转换器
// 提供将 Unix 时间戳转换为时间对象的方法
var NS = &ns把时间戳转为时间{}

// ns把时间戳转为时间 converts Unix timestamps to time objects in the given zone
// The zone is a must, like the rest of the package: the same timestamp gives a different date in each zone
//
// ns把时间戳转为时间 将 Unix 时间戳转换为指定时区的时间对象
// 跟包里其它函数一样必须传时区：同一个时间戳在不同时区里日期可能不同
type ns把时间戳转为时间 struct{}

// Get纳秒时间 converts a Unix timestamp in nanoseconds to a time in the zone
//
// Get纳秒时间 将纳秒级 Unix 时间戳转换为该时区的时间
func (ns *ns把时间戳转为时间) Get纳秒时间(ns纳秒数 int64, z时区位置 *time.Location) time.Time {
	return time.Unix(0, ns纳秒数).In(z时区位置)
}

// Get微秒时间 converts a Unix timestamp in microseconds to a time in the zone
//
// Get微秒时间 将微秒级 Unix 时间戳转换为该时区的时间
func (ns *ns把时间戳转为时间) Get微秒时间(us微秒数 int64, z时区位置 *time.Location) time.Time {
	return time.UnixMicro(us微秒数).In(z时区位置)
}

// Get毫秒时间 converts a Unix timestamp in milliseconds to a time in the zone
//
// Get毫秒时间 将毫秒级 Unix 时间戳转换为该时区的时间
func (ns *ns把时间戳转为时间) Get毫秒时间(ms毫秒数 int64, z时区位置 *time.Location) time.Time {
	return time.UnixMilli(ms毫秒数).In(z时区位置)
}

// Get秒数时间 converts a Unix timestamp in seconds to a time in the zone
//
// Get秒数时间 将秒级 Unix 时间戳转换为该时区的时间
func (ns *ns把时间戳转为时间) Get秒数时间(sn总秒数 int64, z时区位置 *time.Location) time.Time {
	return time.Unix(sn总秒数, 0).In(z时区位置)
}
