package timezh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_wk时间和星期_Is工作日(t *testing.T) {
	require.True(t, timezh.WK.Is工作日(time.Date(2026, 11, 4, 12, 0, 0, 0, timezh.TZ.CST)))  // Wed // 周三
	require.False(t, timezh.WK.Is工作日(time.Date(2026, 11, 7, 12, 0, 0, 0, timezh.TZ.CST))) // Sat // 周六
}

func Test_wk时间和星期_Is休息日(t *testing.T) {
	require.True(t, timezh.WK.Is休息日(time.Date(2026, 11, 1, 12, 0, 0, 0, timezh.TZ.CST)))  // Sun // 周日
	require.False(t, timezh.WK.Is休息日(time.Date(2026, 11, 2, 12, 0, 0, 0, timezh.TZ.CST))) // Mon // 周一
}
