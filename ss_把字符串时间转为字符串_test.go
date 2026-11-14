package timezh_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_ss把字符串时间转为字符串_Get拆分日期时刻(t *testing.T) {
	s日期, s时刻, err := timezh.SS.Get拆分日期时刻("2025-05-15 12:34:56")
	require.NoError(t, err)
	require.Equal(t, "2025-05-15", s日期)
	require.Equal(t, "12:34:56", s时刻)

	_, _, err = timezh.SS.Get拆分日期时刻("2025-05-15")
	require.Error(t, err)
}

func Test_ss把字符串时间转为字符串_Get标准日期(t *testing.T) {
	s日期, err := timezh.SS.Get标准日期("20250515")
	require.NoError(t, err)
	require.Equal(t, "2025-05-15", s日期)

	_, err = timezh.SS.Get标准日期("20251399")
	require.Error(t, err)
}

func Test_ss把字符串时间转为字符串_Get八位数字日期(t *testing.T) {
	s日期, err := timezh.SS.Get八位数字日期("2025-05-15")
	require.NoError(t, err)
	require.Equal(t, "20250515", s日期)

	_, err = timezh.SS.Get八位数字日期("2025-5-15")
	require.Error(t, err)
}

func Test_ss把字符串时间转为字符串_Get年份(t *testing.T) {
	s年份, err := timezh.SS.Get年份("2025-05-15")
	require.NoError(t, err)
	require.Equal(t, "2025", s年份)
}

func Test_ss把字符串时间转为字符串_Get日期月日(t *testing.T) {
	s月日, err := timezh.SS.Get日期月日("2025-05-15")
	require.NoError(t, err)
	require.Equal(t, "05-15", s月日)
}

func Test_ss把字符串时间转为字符串_Get标准时刻(t *testing.T) {
	s时刻, err := timezh.SS.Get标准时刻("123456")
	require.NoError(t, err)
	require.Equal(t, "12:34:56", s时刻)

	_, err = timezh.SS.Get标准时刻("126099")
	require.Error(t, err)
}

func Test_ss把字符串时间转为字符串_Get八位数字转标准时刻(t *testing.T) {
	s时刻, err := timezh.SS.Get八位数字转标准时刻("12345699")
	require.NoError(t, err)
	require.Equal(t, "12:34:56", s时刻)
}

func Test_ss把字符串时间转为字符串_Get六位数字时刻(t *testing.T) {
	s时刻, err := timezh.SS.Get六位数字时刻("12:34:56")
	require.NoError(t, err)
	require.Equal(t, "123456", s时刻)

	_, err = timezh.SS.Get六位数字时刻("12:34")
	require.Error(t, err)
}

func Test_ss把字符串时间转为字符串_Get标准时分(t *testing.T) {
	s时分, err := timezh.SS.Get标准时分("1530")
	require.NoError(t, err)
	require.Equal(t, "15:30", s时分)

	_, err = timezh.SS.Get标准时分("12")
	require.Error(t, err)
	_, err = timezh.SS.Get标准时分("ab30")
	require.Error(t, err)
}

func Test_ss把字符串时间转为字符串_Get取出日期(t *testing.T) {
	s日期, err := timezh.SS.Get取出日期("2025-05-15 12:34:56")
	require.NoError(t, err)
	require.Equal(t, "2025-05-15", s日期)
}

func Test_ss把字符串时间转为字符串_Get取出时刻(t *testing.T) {
	s时刻, err := timezh.SS.Get取出时刻("2025-05-15 12:34:56")
	require.NoError(t, err)
	require.Equal(t, "12:34:56", s时刻)

	_, err = timezh.SS.Get取出时刻("2025-05-15")
	require.Error(t, err)
}
