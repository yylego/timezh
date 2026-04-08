package timezh_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_ss把字符串时间转为字符串_Get把标准时刻转换为6位数时刻(t *testing.T) {
	HMS6n, err := timezh.SS.Get把标准时刻转换为6位数时刻("12:34:56")
	require.NoError(t, err)
	t.Log(HMS6n)
	HcMcS, err := timezh.SS.Get把6位数时刻转换为标准时刻(HMS6n)
	require.NoError(t, err)
	t.Log(HcMcS)
}

func Test_ss把字符串时间转为字符串_Get把标准时刻转换为6位数时刻_错误格式(t *testing.T) {
	_, err := timezh.SS.Get把标准时刻转换为6位数时刻("abc")
	require.Error(t, err)
	t.Log("expected:", err)
}

func Test_ss把字符串时间转为字符串_Get把4位时分转换为时和杠和分(t *testing.T) {
	v, err := timezh.SS.Get把4位时分转换为时和杠和分("1530")
	require.NoError(t, err)
	require.Equal(t, "15:30", v)
}

func Test_ss把字符串时间转为字符串_Get把4位时分转换为时和杠和分_错误格式(t *testing.T) {
	_, err := timezh.SS.Get把4位时分转换为时和杠和分("12")
	require.Error(t, err)
	t.Log("expected:", err)
}

func Test_ss把字符串时间转为字符串_Get把年月日和时分秒分开返回(t *testing.T) {
	日期, 时间, err := timezh.SS.Get把年月日和时分秒分开返回("2025-05-15 12:34:56")
	require.NoError(t, err)
	require.Equal(t, "2025-05-15", 日期)
	require.Equal(t, "12:34:56", 时间)
}

func Test_ss把字符串时间转为字符串_Get把年月日和时分秒分开返回_错误格式(t *testing.T) {
	_, _, err := timezh.SS.Get把年月日和时分秒分开返回("not-datetime")
	require.Error(t, err)
	t.Log("expected:", err)
}
