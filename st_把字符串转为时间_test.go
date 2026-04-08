package timezh_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_st把字符串转为时间_Get时间(t *testing.T) {
	{
		v时间, err := timezh.ST.Get时间("2025-05-15 15:34:06", timezh.TZ.ICT)
		require.NoError(t, err)
		t.Log(v时间.UTC())
	}
	{
		v时间, err := timezh.ST.Get时间("2025-05-15 15:34:06", timezh.TZ.CST)
		require.NoError(t, err)
		t.Log(v时间.UTC())
	}
	{
		v时间, err := timezh.ST.Get时间("2025-05-15 15:34:06", timezh.TZ.UTC)
		require.NoError(t, err)
		t.Log(v时间.UTC())
	}
}

func Test_st把字符串转为时间_Get时间_错误格式(t *testing.T) {
	_, err := timezh.ST.Get时间("not-a-time", timezh.TZ.CST)
	require.Error(t, err)
	t.Log("expected:", err)
}

func Test_st把字符串转为时间_Get日期和分时(t *testing.T) {
	v时间, err := timezh.ST.Get日期和分时("2025-05-15 09:30", timezh.TZ.CST)
	require.NoError(t, err)
	t.Log(v时间)
}

func Test_st把字符串转为时间_Get日期和分时_错误格式(t *testing.T) {
	_, err := timezh.ST.Get日期和分时("abc", timezh.TZ.CST)
	require.Error(t, err)
	t.Log("expected:", err)
}

func Test_st把字符串转为时间_Get通过日期得到时间(t *testing.T) {
	v时间, err := timezh.ST.Get通过日期得到时间("2025-05-15", timezh.TZ.CST)
	require.NoError(t, err)
	t.Log(v时间)
}

func Test_st把字符串转为时间_Get通过日期得到时间_错误格式(t *testing.T) {
	_, err := timezh.ST.Get通过日期得到时间("20250515", timezh.TZ.CST)
	require.Error(t, err)
	t.Log("expected:", err)
}

func Test_st把字符串转为时间_Get通过时刻得到当日时间(t *testing.T) {
	v时间, err := timezh.ST.Get通过时刻得到当日时间("15:03:22", timezh.TZ.CST)
	require.NoError(t, err)
	t.Log(v时间)
}

func Test_st把字符串转为时间_Get通过时刻得到当日时间_错误格式(t *testing.T) {
	_, err := timezh.ST.Get通过时刻得到当日时间("25:99:99", timezh.TZ.CST)
	require.Error(t, err)
	t.Log("expected:", err)
}
