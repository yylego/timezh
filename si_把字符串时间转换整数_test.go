package timezh_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_si把字符串时间转换整数_Get通过日期得到年份(t *testing.T) {
	v年份, err := timezh.SI.Get通过日期得到年份("2021-01-23")
	require.NoError(t, err)
	require.Equal(t, 2021, v年份)
	t.Log(v年份)
}

func Test_si把字符串时间转换整数_Get通过日期得到年份_错误格式(t *testing.T) {
	_, err := timezh.SI.Get通过日期得到年份("20210123")
	require.Error(t, err)
	t.Log("expected:", err)
}

func Test_si把字符串时间转换整数_Get通过时刻得到几时几分秒数(t *testing.T) {
	V几时, v几分, v秒数, err := timezh.SI.Get通过时刻得到几时几分秒数("16:44:02")
	require.NoError(t, err)
	require.Equal(t, 16, V几时)
	require.Equal(t, 44, v几分)
	require.Equal(t, 2, v秒数)
	t.Log(V几时, v几分, v秒数)
}

func Test_si把字符串时间转换整数_Get通过时刻得到几时几分秒数_错误格式(t *testing.T) {
	_, _, _, err := timezh.SI.Get通过时刻得到几时几分秒数("abc")
	require.Error(t, err)
	t.Log("expected:", err)
}

func Test_si把字符串时间转换整数_Get通过日期得到年份月份哪日(t *testing.T) {
	v年份, v月份, v几号, err := timezh.SI.Get通过日期得到年份月份哪日("2021-01-23")
	require.NoError(t, err)
	require.Equal(t, 2021, v年份)
	require.Equal(t, 1, v月份)
	require.Equal(t, 23, v几号)
	t.Log(v年份, v月份, v几号)
}

func Test_si把字符串时间转换整数_Get通过日期得到年份月份哪日_错误格式(t *testing.T) {
	_, _, _, err := timezh.SI.Get通过日期得到年份月份哪日("not-a-date")
	require.Error(t, err)
	t.Log("expected:", err)
}

func Test_si把字符串时间转换整数_Get通过时间得到整数(t *testing.T) {
	v年份, v月份, v几号, V几时, v几分, v秒数, err := timezh.SI.Get通过时间得到整数("2021-01-23 16:47:03")
	require.NoError(t, err)
	require.Equal(t, 2021, v年份)
	require.Equal(t, 1, v月份)
	require.Equal(t, 23, v几号)
	require.Equal(t, 16, V几时)
	require.Equal(t, 47, v几分)
	require.Equal(t, 3, v秒数)
	t.Log(v年份, v月份, v几号, V几时, v几分, v秒数)
}

func Test_si把字符串时间转换整数_Get通过时间得到整数_错误格式(t *testing.T) {
	_, _, _, _, _, _, err := timezh.SI.Get通过时间得到整数("bad-time")
	require.Error(t, err)
	t.Log("expected:", err)
}
