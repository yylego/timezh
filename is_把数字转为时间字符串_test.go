package timezh_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_is把数字转为时间字符串_Get日期(t *testing.T) {
	s日期, err := timezh.IS.Get日期(2021, 1, 23)
	require.NoError(t, err)
	require.Equal(t, "2021-01-23", s日期)

	_, err = timezh.IS.Get日期(2025, 13, 1)
	require.Error(t, err)
	_, err = timezh.IS.Get日期(2025, 2, 29)
	require.Error(t, err)
}

func Test_is把数字转为时间字符串_Get时刻(t *testing.T) {
	s时刻, err := timezh.IS.Get时刻(9, 5, 3)
	require.NoError(t, err)
	require.Equal(t, "09:05:03", s时刻)

	_, err = timezh.IS.Get时刻(24, 0, 0)
	require.Error(t, err)
}

func Test_is把数字转为时间字符串_Get时间(t *testing.T) {
	s时间, err := timezh.IS.Get时间(2021, 1, 23, 9, 5, 3)
	require.NoError(t, err)
	require.Equal(t, "2021-01-23 09:05:03", s时间)

	_, err = timezh.IS.Get时间(2021, 1, 23, 9, 60, 3)
	require.Error(t, err)
}
