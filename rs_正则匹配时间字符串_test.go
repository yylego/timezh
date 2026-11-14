package timezh_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_rs正则匹配时间字符串_Match日期(t *testing.T) {
	require.True(t, timezh.RS.Match日期("2018-01-01"))
	require.True(t, timezh.RS.Match日期("2024-02-29")) // leap date // 闰年 2 月 29 日
	require.False(t, timezh.RS.Match日期("2018-01-010"))
	require.False(t, timezh.RS.Match日期("2018-1-01"))
	require.False(t, timezh.RS.Match日期("2018-01-01 2018-01-01"))
	// The shape is right, but no such date
	// 格式对，但没有这一天
	require.False(t, timezh.RS.Match日期("2025-13-01"))
	require.False(t, timezh.RS.Match日期("2025-02-29"))
}

func Test_rs正则匹配时间字符串_Match时刻(t *testing.T) {
	require.True(t, timezh.RS.Match时刻("09:00:01"))
	require.True(t, timezh.RS.Match时刻("23:59:59"))
	require.False(t, timezh.RS.Match时刻("9:00:01"))
	require.False(t, timezh.RS.Match时刻("24:00:00"))
	require.False(t, timezh.RS.Match时刻("10:60:00"))
}

func Test_rs正则匹配时间字符串_Match时间(t *testing.T) {
	require.True(t, timezh.RS.Match时间("2018-01-01 10:00:30"))
	require.False(t, timezh.RS.Match时间("2018-01-01"))
	require.False(t, timezh.RS.Match时间("2018-01-01 9:00:30"))
	require.False(t, timezh.RS.Match时间("2025-02-30 10:00:30"))
}

func Test_rs正则匹配时间字符串_Match日期分时(t *testing.T) {
	require.True(t, timezh.RS.Match日期分时("2025-05-15 09:30"))
	require.False(t, timezh.RS.Match日期分时("2025-05-15 9:30"))
	require.False(t, timezh.RS.Match日期分时("2025-05-15 09:30:00"))
}

func Test_rs正则匹配时间字符串_Match日期零点(t *testing.T) {
	require.True(t, timezh.RS.Match日期零点("2025-05-15 00:00:00"))
	require.False(t, timezh.RS.Match日期零点("2025-05-15 00:00:01"))
	require.False(t, timezh.RS.Match日期零点("2025-05-32 00:00:00"))
}

func Test_rs正则匹配时间字符串_Match八位数字日期(t *testing.T) {
	require.True(t, timezh.RS.Match八位数字日期("20250515"))
	require.False(t, timezh.RS.Match八位数字日期("2025515"))
	require.False(t, timezh.RS.Match八位数字日期("20251399"))
}

func Test_rs正则匹配时间字符串_Match六位数字时刻(t *testing.T) {
	require.True(t, timezh.RS.Match六位数字时刻("150405"))
	require.False(t, timezh.RS.Match六位数字时刻("15040"))
	require.False(t, timezh.RS.Match六位数字时刻("256161"))
}

func Test_rs正则匹配时间字符串_Match八位数字时刻(t *testing.T) {
	require.True(t, timezh.RS.Match八位数字时刻("15040599"))
	require.False(t, timezh.RS.Match八位数字时刻("1504059"))
	require.False(t, timezh.RS.Match八位数字时刻("25040599"))
}

func Test_rs正则匹配时间字符串_Match四位数字时分(t *testing.T) {
	require.True(t, timezh.RS.Match四位数字时分("1530"))
	require.False(t, timezh.RS.Match四位数字时分("153"))
	require.False(t, timezh.RS.Match四位数字时分("2460"))
}

func Test_rs正则匹配时间字符串_Match年份(t *testing.T) {
	require.True(t, timezh.RS.Match年份("2025"))
	require.False(t, timezh.RS.Match年份("25"))
}
