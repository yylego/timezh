package timezh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_nt把数字转化为时间_New当日时刻时间(t *testing.T) {
	v时间, err := timezh.NT.New当日时刻时间(3, 3, 3, timezh.TZ.CST)
	require.NoError(t, err)
	v当天 := time.Now().In(timezh.TZ.CST)
	require.Equal(t, time.Date(v当天.Year(), v当天.Month(), v当天.Day(), 3, 3, 3, 0, timezh.TZ.CST), v时间)

	// 25 o'clock gives an error, it does not move on to the next date
	// 25 点不会顺延到第二天
	_, err = timezh.NT.New当日时刻时间(25, 0, 0, timezh.TZ.CST)
	require.Error(t, err)
}

func Test_nt把数字转化为时间_New某日零点时间(t *testing.T) {
	v时间, err := timezh.NT.New某日零点时间(2025, 5, 15, timezh.TZ.CST)
	require.NoError(t, err)
	require.Equal(t, time.Date(2025, 5, 15, 0, 0, 0, 0, timezh.TZ.CST), v时间)

	// 02-30 gives an error, it does not move on to 03-02
	// 2 月 30 日不会顺延成 3 月 2 日
	_, err = timezh.NT.New某日零点时间(2025, 2, 30, timezh.TZ.CST)
	require.Error(t, err)
}

func Test_nt把数字转化为时间_New当日零点时间(t *testing.T) {
	v时间 := timezh.NT.New当日零点时间(time.Date(2025, 5, 16, 19, 30, 45, 123, timezh.TZ.CST))
	require.Equal(t, time.Date(2025, 5, 16, 0, 0, 0, 0, timezh.TZ.CST), v时间)
}
