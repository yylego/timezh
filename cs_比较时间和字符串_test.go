package timezh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_cs比较时间和字符串_Is时刻之间(t *testing.T) {
	v时间 := time.Date(2026, 11, 4, 10, 30, 0, 0, timezh.TZ.CST)
	for _, item := range []struct {
		s起始时刻 string
		e结束时刻 string
		want  bool
	}{
		{"09:15:00", "15:00:00", true},
		{"10:30:00", "10:30:00", true},
		{"11:00:00", "15:00:00", false},
	} {
		match, err := timezh.CS.Is时刻之间(v时间, item.s起始时刻, item.e结束时刻)
		require.NoError(t, err)
		require.Equal(t, item.want, match)
	}

	// "9:00:00" sorts past "10:30:00" as a string, so a shape that is not zero-padded gives an error instead of a wrong result
	// "9:00:00" 按字符串比会排到 "10:30:00" 后面，所以没补零的格式返回错误，不给出错的结果
	_, err := timezh.CS.Is时刻之间(v时间, "9:00:00", "17:00:00")
	require.Error(t, err)
}

func Test_cs比较时间和字符串_Is晚于时刻(t *testing.T) {
	v时间 := time.Date(2026, 11, 4, 10, 30, 0, 0, timezh.TZ.CST)
	match, err := timezh.CS.Is晚于时刻(v时间, "10:00:00")
	require.NoError(t, err)
	require.True(t, match)

	match, err = timezh.CS.Is晚于时刻(v时间, "11:00:00")
	require.NoError(t, err)
	require.False(t, match)
}

func Test_cs比较时间和字符串_Is早于时刻(t *testing.T) {
	v时间 := time.Date(2026, 11, 4, 10, 30, 0, 0, timezh.TZ.CST)
	match, err := timezh.CS.Is早于时刻(v时间, "11:00:00")
	require.NoError(t, err)
	require.True(t, match)

	match, err = timezh.CS.Is早于时刻(v时间, "10:00:00")
	require.NoError(t, err)
	require.False(t, match)

	_, err = timezh.CS.Is早于时刻(v时间, "9:00")
	require.Error(t, err)
}
