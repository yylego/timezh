package timezh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_sd比较字符串得到间隔_Get距今日期间隔(t *testing.T) {
	v间隔, err := timezh.SD.Get距今日期间隔("2099-12-31", timezh.TZ.CST)
	require.NoError(t, err)
	require.Greater(t, v间隔, time.Duration(0))
}

func Test_sd比较字符串得到间隔_Get起点日期间隔(t *testing.T) {
	v时间 := time.Date(2025, 5, 15, 12, 0, 0, 0, timezh.TZ.CST)
	v间隔, err := timezh.SD.Get起点日期间隔(v时间, "2025-05-16")
	require.NoError(t, err)
	require.Equal(t, 12*time.Hour, v间隔)
}

func Test_sd比较字符串得到间隔_Get距今时刻间隔(t *testing.T) {
	v间隔, err := timezh.SD.Get距今时刻间隔("00:00:00", timezh.TZ.CST)
	require.NoError(t, err)
	require.LessOrEqual(t, v间隔, time.Duration(0))

	_, err = timezh.SD.Get距今时刻间隔("bad", timezh.TZ.CST)
	require.Error(t, err)
}

func Test_sd比较字符串得到间隔_Get起点时刻间隔(t *testing.T) {
	v时间 := time.Date(2025, 5, 15, 12, 0, 0, 0, timezh.TZ.CST)
	v间隔, err := timezh.SD.Get起点时刻间隔(v时间, "13:30:00")
	require.NoError(t, err)
	require.Equal(t, 90*time.Minute, v间隔)

	// A minute of 99 is not a time-of-moment, it does not move on to the next 60 minutes
	// 99 分不是合法时刻，不会滚到下一个小时去
	_, err = timezh.SD.Get起点时刻间隔(v时间, "13:99:00")
	require.Error(t, err)
}

func Test_sd比较字符串得到间隔_Get距今时间间隔(t *testing.T) {
	v间隔, err := timezh.SD.Get距今时间间隔("2021-01-23 17:21:00", timezh.TZ.CST)
	require.NoError(t, err)
	require.Less(t, v间隔, time.Duration(0))
}

func Test_sd比较字符串得到间隔_Get起点时间间隔(t *testing.T) {
	v时间 := time.Date(2025, 5, 15, 12, 0, 0, 0, timezh.TZ.CST)
	v间隔, err := timezh.SD.Get起点时间间隔(v时间, "2025-05-15 12:00:30")
	require.NoError(t, err)
	require.Equal(t, 30*time.Second, v间隔)
}

func Test_sd比较字符串得到间隔_Get两个时间间隔(t *testing.T) {
	v间隔, err := timezh.SD.Get两个时间间隔("2021-01-23 17:21:00", "2021-01-24 18:22:01")
	require.NoError(t, err)
	require.Equal(t, 25*time.Hour+time.Minute+time.Second, v间隔)

	_, err = timezh.SD.Get两个时间间隔("bad-time", "2021-02-24 18:22:01")
	require.Error(t, err)
}
