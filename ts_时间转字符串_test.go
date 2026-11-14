package timezh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_ts日期转字符串_Get截至当日工作日期(t *testing.T) {
	workdates := timezh.TS.D日期.Get截至当日工作日期(10, timezh.TZ.CST)
	require.Len(t, workdates, 10)
	for _, s日期 := range workdates {
		v时间, err := timezh.TS.D日期.Get转为时间(s日期, timezh.TZ.CST)
		require.NoError(t, err)
		require.True(t, timezh.WK.Is工作日(v时间), s日期)
	}
}

func Test_ts日期转字符串_Get截至某日工作日期(t *testing.T) {
	// 2026-11-02 is a Mon, so the next one back skips Sat and Sun and lands on Fri 10-30
	// 2026-11-02 是周一，再往前一个要跳过周末，落到周五 10-30
	v周一 := time.Date(2026, 11, 2, 9, 0, 0, 0, timezh.TZ.CST)
	require.Equal(t, []string{"2026-11-02", "2026-10-30"}, timezh.TS.D日期.Get截至某日工作日期(v周一, 2))
	require.Empty(t, timezh.TS.D日期.Get截至某日工作日期(v周一, 0))
}

func Test_ts时间转字符串_Get现在时间(t *testing.T) {
	require.True(t, timezh.RS.Match时间(timezh.TS.T时间.Get现在时间(timezh.TZ.CST)))
}

func Test_ts时间转字符串_Get转为文本(t *testing.T) {
	require.Equal(t, "15:04:05", timezh.TS.T时刻.Get转为文本(time.Date(2026, 11, 4, 15, 4, 5, 0, timezh.TZ.CST)))
}

func Test_ts日期转字符串_Get转为文本(t *testing.T) {
	require.Equal(t, "2026-11-04", timezh.TS.D日期.Get转为文本(time.Date(2026, 11, 4, 15, 4, 5, 0, timezh.TZ.CST)))
}

func Test_ts日期转字符串_Get月初日期(t *testing.T) {
	v时间 := time.Date(2025, 5, 16, 9, 0, 0, 0, timezh.TZ.CST)
	require.Equal(t, "2025-05-01", timezh.TS.D日期.Get月初日期(v时间))
	// Each component keeps its own layout
	// 每个格式化器按自己的布局输出
	require.Equal(t, "20250501", timezh.TS.D八位数字日期.Get月初日期(v时间))
}

func Test_ts日期转字符串_Get当日日期(t *testing.T) {
	require.Equal(t, time.Now().In(timezh.TZ.CST).Format(time.DateOnly), timezh.TS.D日期.Get当日日期(timezh.TZ.CST))
}

func Test_ts日期转字符串_Get昨日日期(t *testing.T) {
	require.Equal(t, time.Now().In(timezh.TZ.CST).AddDate(0, 0, -1).Format(time.DateOnly), timezh.TS.D日期.Get昨日日期(timezh.TZ.CST))
}

func Test_ts日期转字符串_Get转为时间(t *testing.T) {
	v时间, err := timezh.TS.D八位数字日期.Get转为时间("20250516", timezh.TZ.CST)
	require.NoError(t, err)
	require.Equal(t, time.Date(2025, 5, 16, 0, 0, 0, 0, timezh.TZ.CST), v时间)

	_, err = timezh.TS.D八位数字日期.Get转为时间("2025-05-16", timezh.TZ.CST)
	require.Error(t, err)
	_, err = timezh.TS.D八位数字日期.Get转为时间("20250230", timezh.TZ.CST)
	require.Error(t, err)
}

func Test_ts日期转字符串_Get明日日期(t *testing.T) {
	require.Equal(t, time.Now().In(timezh.TZ.CST).AddDate(0, 0, 1).Format(time.DateOnly), timezh.TS.D日期.Get明日日期(timezh.TZ.CST))
}

func Test_ts时间转字符串_Get转为时间_时分(t *testing.T) {
	require.Equal(t, "09:30", timezh.TS.T时分.Get转为文本(time.Date(2025, 5, 16, 9, 30, 45, 0, timezh.TZ.CST)))
	require.True(t, timezh.TS.T时分.Match格式("23:59"))
	require.False(t, timezh.TS.T时分.Match格式("9:30"))
	require.False(t, timezh.TS.T时分.Match格式("24:00"))
}

func TestTS_各个字段自洽(t *testing.T) {
	// Each field of this time has two digits and none equals the reference time, so a layout part that is not fixed-width would show up
	// 这个时间的每个字段都是两位数、又都跟参考时间不同，布局里位数不固定的部分会在这里露出来
	v时间 := time.Date(2025, 12, 31, 23, 59, 59, 0, timezh.TZ.CST)
	for _, field := range []interface {
		Get转为文本(time.Time) string
		Match格式(string) bool
	}{
		timezh.TS.D日期, timezh.TS.Y年份, timezh.TS.D月日, timezh.TS.D八位数字日期,
		timezh.TS.T时刻, timezh.TS.T时分, timezh.TS.T时间, timezh.TS.T日期分时,
		timezh.TS.T国际格式时间, timezh.TS.T十四位数字时间, timezh.TS.T六位数字时刻,
	} {
		s文本 := field.Get转为文本(v时间)
		require.True(t, field.Match格式(s文本), s文本)
	}
}

func Test_ts日期转字符串_Match格式(t *testing.T) {
	require.True(t, timezh.TS.D日期.Match格式("2025-05-16"))
	require.False(t, timezh.TS.D日期.Match格式("2025-5-16"))
	require.True(t, timezh.TS.D月日.Match格式("05-16"))
	require.False(t, timezh.TS.D月日.Match格式("13-16"))
}

func Test_ts日期转字符串_Get加减天数日期(t *testing.T) {
	s日期, err := timezh.TS.D日期.Get加减天数日期("2025-03-01", -1)
	require.NoError(t, err)
	require.Equal(t, "2025-02-28", s日期)

	s日期, err = timezh.TS.D八位数字日期.Get加减天数日期("20241231", 1)
	require.NoError(t, err)
	require.Equal(t, "20250101", s日期)

	_, err = timezh.TS.D日期.Get加减天数日期("2025-02-30", 1)
	require.Error(t, err)
}

func Test_ts日期转字符串_Get区间每天日期(t *testing.T) {
	dates, err := timezh.TS.D日期.Get区间每天日期("2024-02-27", "2024-03-01")
	require.NoError(t, err)
	require.Equal(t, []string{"2024-02-27", "2024-02-28", "2024-02-29", "2024-03-01"}, dates)

	dates, err = timezh.TS.D日期.Get区间每天日期("2024-03-01", "2024-02-27")
	require.NoError(t, err)
	require.Empty(t, dates)

	_, err = timezh.TS.D日期.Get区间每天日期("2024-03-01", "bad")
	require.Error(t, err)
}

func Test_ts日期转字符串_Get月末日期(t *testing.T) {
	require.Equal(t, "2024-02-29", timezh.TS.D日期.Get月末日期(time.Date(2024, 2, 10, 9, 0, 0, 0, timezh.TZ.CST)))
	require.Equal(t, "20251231", timezh.TS.D八位数字日期.Get月末日期(time.Date(2025, 12, 31, 23, 0, 0, 0, timezh.TZ.CST)))
}

func Test_ts日期转字符串_Get某月月初日期(t *testing.T) {
	v时间 := time.Date(2025, 3, 31, 9, 0, 0, 0, timezh.TZ.CST)
	require.Equal(t, "2025-02-01", timezh.TS.D日期.Get某月月初日期(v时间, -1))
	require.Equal(t, "2024-12-01", timezh.TS.D日期.Get某月月初日期(v时间, -3))
	require.Equal(t, "2025-04-01", timezh.TS.D日期.Get某月月初日期(v时间, 1))
}

func Test_ts日期转字符串_Get某月月末日期(t *testing.T) {
	v时间 := time.Date(2025, 3, 31, 9, 0, 0, 0, timezh.TZ.CST)
	require.Equal(t, "2025-02-28", timezh.TS.D日期.Get某月月末日期(v时间, -1))
	require.Equal(t, "2024-02-29", timezh.TS.D日期.Get某月月末日期(v时间, -13))
	require.Equal(t, "2025-04-30", timezh.TS.D日期.Get某月月末日期(v时间, 1))
}

func Test_ts日期转字符串_Get当周周一日期(t *testing.T) {
	// 2026-11-04 is a Wed, 2026-11-08 is a Sun: both are in the week from Mon 11-02
	// 2026-11-04 是周三，2026-11-08 是周日：都在 11-02 周一开始的那一周
	require.Equal(t, "2026-11-02", timezh.TS.D日期.Get当周周一日期(time.Date(2026, 11, 4, 9, 0, 0, 0, timezh.TZ.CST)))
	require.Equal(t, "2026-11-02", timezh.TS.D日期.Get当周周一日期(time.Date(2026, 11, 8, 9, 0, 0, 0, timezh.TZ.CST)))
	require.Equal(t, "2026-11-02", timezh.TS.D日期.Get当周周一日期(time.Date(2026, 11, 2, 9, 0, 0, 0, timezh.TZ.CST)))
}

func Test_ts日期转字符串_Get当周周日日期(t *testing.T) {
	require.Equal(t, "2026-11-08", timezh.TS.D日期.Get当周周日日期(time.Date(2026, 11, 4, 9, 0, 0, 0, timezh.TZ.CST)))
	require.Equal(t, "2026-11-08", timezh.TS.D日期.Get当周周日日期(time.Date(2026, 11, 8, 9, 0, 0, 0, timezh.TZ.CST)))
}

func Test_ts时间转字符串_Get转为时间(t *testing.T) {
	v时间, err := timezh.TS.T十四位数字时间.Get转为时间("20250516193045", timezh.TZ.CST)
	require.NoError(t, err)
	require.Equal(t, time.Date(2025, 5, 16, 19, 30, 45, 0, timezh.TZ.CST), v时间)

	v时间, err = timezh.TS.T国际格式时间.Get转为时间("2025-05-16T19:30:45", timezh.TZ.CST)
	require.NoError(t, err)
	require.Equal(t, time.Date(2025, 5, 16, 19, 30, 45, 0, timezh.TZ.CST), v时间)

	_, err = timezh.TS.T十四位数字时间.Get转为时间("20250516253045", timezh.TZ.CST)
	require.Error(t, err)
}

func Test_ts时间转字符串_Match格式(t *testing.T) {
	require.True(t, timezh.TS.T六位数字时刻.Match格式("193045"))
	require.False(t, timezh.TS.T六位数字时刻.Match格式("19:30:45"))
}

func Test_ts时间转字符串_Get转为时间_各个时区(t *testing.T) {
	// The same clock time read in each zone gives a different instant
	// 同一个钟点放到不同时区，得到的是不同的时刻
	for _, item := range []struct {
		zone *time.Location
		want string
	}{
		{timezh.TZ.CST, "2025-05-15T07:34:06Z"},
		{timezh.TZ.ICT, "2025-05-15T08:34:06Z"},
		{timezh.TZ.UTC, "2025-05-15T15:34:06Z"},
	} {
		v时间, err := timezh.TS.T时间.Get转为时间("2025-05-15 15:34:06", item.zone)
		require.NoError(t, err)
		require.Equal(t, item.want, v时间.UTC().Format(time.RFC3339))
	}

	_, err := timezh.TS.T时间.Get转为时间("2025-05-15 9:34:06", timezh.TZ.CST)
	require.Error(t, err)
}

func Test_ts时间转字符串_Get转为时间_日期分时(t *testing.T) {
	v时间, err := timezh.TS.T日期分时.Get转为时间("2025-05-15 09:30", timezh.TZ.CST)
	require.NoError(t, err)
	require.Equal(t, time.Date(2025, 5, 15, 9, 30, 0, 0, timezh.TZ.CST), v时间)

	_, err = timezh.TS.T日期分时.Get转为时间("2025-05-15 9:30", timezh.TZ.CST)
	require.Error(t, err)
}
