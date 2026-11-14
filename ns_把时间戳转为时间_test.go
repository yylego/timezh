package timezh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

// v时间戳时刻 is 2025-05-15 16:30:00 UTC, which is 2025-05-16 00:30:00 at UTC+8
// v时间戳时刻 是 UTC 的 2025-05-15 16:30:00，在 UTC+8 已经是 2025-05-16 00:30:00
var v时间戳时刻 = time.Date(2025, 5, 15, 16, 30, 0, 0, time.UTC)

func Test_ns把时间戳转为时间_Get纳秒时间(t *testing.T) {
	v时间 := timezh.NS.Get纳秒时间(v时间戳时刻.UnixNano(), timezh.TZ.CST)
	require.Equal(t, "2025-05-16 00:30:00", timezh.TS.T时间.Get转为文本(v时间))
	require.Equal(t, time.Date(1969, 12, 31, 23, 59, 59, 500000000, time.UTC), timezh.NS.Get纳秒时间(-5e8, timezh.TZ.UTC))
}

func Test_ns把时间戳转为时间_Get微秒时间(t *testing.T) {
	v时间 := timezh.NS.Get微秒时间(v时间戳时刻.UnixMicro(), timezh.TZ.CST)
	require.Equal(t, "2025-05-16 00:30:00", timezh.TS.T时间.Get转为文本(v时间))
}

func Test_ns把时间戳转为时间_Get毫秒时间(t *testing.T) {
	v时间 := timezh.NS.Get毫秒时间(v时间戳时刻.UnixMilli(), timezh.TZ.CST)
	require.Equal(t, "2025-05-16 00:30:00", timezh.TS.T时间.Get转为文本(v时间))
}

func Test_ns把时间戳转为时间_Get秒数时间(t *testing.T) {
	// The same timestamp gives a different date in each zone
	// 同一个时间戳，在不同时区里是不同的日期
	require.Equal(t, "2025-05-16", timezh.TS.D日期.Get转为文本(timezh.NS.Get秒数时间(v时间戳时刻.Unix(), timezh.TZ.CST)))
	require.Equal(t, "2025-05-15", timezh.TS.D日期.Get转为文本(timezh.NS.Get秒数时间(v时间戳时刻.Unix(), timezh.TZ.UTC)))
}
