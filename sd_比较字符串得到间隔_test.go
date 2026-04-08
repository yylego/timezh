package timezh_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_sd比较字符串得到间隔_Get现在到日期的间隔(t *testing.T) {
	d1, err := timezh.SD.Get现在到日期的间隔("2038-01-01", timezh.TZ.CST)
	require.NoError(t, err)
	t.Log(d1)

	d2, err := timezh.SD.Get现在到日期的间隔("2099-12-31", timezh.TZ.CST)
	require.NoError(t, err)
	t.Log(d2)
}

func Test_sd比较字符串得到间隔_Get现在到当天时刻的间隔(t *testing.T) {
	d1, err := timezh.SD.Get现在到当天时刻的间隔("00:00:00", timezh.TZ.UTC)
	require.NoError(t, err)
	t.Log(d1)

	d2, err := timezh.SD.Get现在到当天时刻的间隔("23:59:59", timezh.TZ.CST)
	require.NoError(t, err)
	t.Log(d2)
}

func Test_sd比较字符串得到间隔_Get现在到时间的间隔(t *testing.T) {
	d1, err := timezh.SD.Get现在到时间的间隔("2021-01-23 17:21:00", timezh.TZ.UTC)
	require.NoError(t, err)
	t.Log(d1)

	d2, err := timezh.SD.Get现在到时间的间隔("2021-01-23 17:21:00", timezh.TZ.CST)
	require.NoError(t, err)
	t.Log(d2)
}

func Test_sd比较字符串得到间隔_Get两字符串时间间隔(t *testing.T) {
	d1, err := timezh.SD.Get两字符串时间间隔("2021-01-23 17:21:00", "2021-02-24 18:22:01")
	require.NoError(t, err)
	t.Log(d1)
}

func Test_sd比较字符串得到间隔_Get两字符串时间间隔_错误格式(t *testing.T) {
	_, err := timezh.SD.Get两字符串时间间隔("bad-time", "2021-02-24 18:22:01")
	require.Error(t, err)
	t.Log("expected:", err)
}

func Test_sd比较字符串得到间隔_Get现在到当天时刻的间隔_错误格式(t *testing.T) {
	_, err := timezh.SD.Get现在到当天时刻的间隔("bad", timezh.TZ.CST)
	require.Error(t, err)
	t.Log("expected:", err)
}
