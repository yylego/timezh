package timezh_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_si把字符串时间转换整数_Get年份(t *testing.T) {
	v年份, err := timezh.SI.Get年份("2021-01-23")
	require.NoError(t, err)
	require.Equal(t, 2021, v年份)

	_, err = timezh.SI.Get年份("20210123")
	require.Error(t, err)
}

func Test_si把字符串时间转换整数_Get拆分日期(t *testing.T) {
	v年份, v月份, v几号, err := timezh.SI.Get拆分日期("2021-01-23")
	require.NoError(t, err)
	require.Equal(t, []int{2021, 1, 23}, []int{v年份, v月份, v几号})

	// Text past the end of the date and dates that do not exist both return errors
	// 日期后面多了字符、日期不存在，都报错
	_, _, _, err = timezh.SI.Get拆分日期("2021-01-23xyz")
	require.Error(t, err)
	_, _, _, err = timezh.SI.Get拆分日期("2021-13-23")
	require.Error(t, err)
}

func Test_si把字符串时间转换整数_Get拆分时刻(t *testing.T) {
	v几时, v几分, v秒数, err := timezh.SI.Get拆分时刻("16:44:02")
	require.NoError(t, err)
	require.Equal(t, []int{16, 44, 2}, []int{v几时, v几分, v秒数})

	_, _, _, err = timezh.SI.Get拆分时刻("1:2:3")
	require.Error(t, err)
	_, _, _, err = timezh.SI.Get拆分时刻("16:44:99")
	require.Error(t, err)
}

func Test_si把字符串时间转换整数_Get拆分时间(t *testing.T) {
	v年份, v月份, v几号, v几时, v几分, v秒数, err := timezh.SI.Get拆分时间("2021-01-23 16:47:03")
	require.NoError(t, err)
	require.Equal(t, []int{2021, 1, 23, 16, 47, 3}, []int{v年份, v月份, v几号, v几时, v几分, v秒数})

	_, _, _, _, _, _, err = timezh.SI.Get拆分时间("2021-01-23 25:47:03")
	require.Error(t, err)
}
