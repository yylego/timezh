package timezh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func TestTimeZone(t *testing.T) {
	// One instant read in each zone: the clock time shows each offset
	// 同一时刻放到各个时区里看：几点钟就是各自的偏移量
	v时刻 := time.Date(2026, 11, 4, 0, 0, 0, 0, time.UTC)
	require.Equal(t, 8, v时刻.In(timezh.TZ.CST).Hour())
	require.Equal(t, 0, v时刻.In(timezh.TZ.UTC).Hour())
	require.Equal(t, 0, v时刻.In(timezh.TZ.GMT).Hour())
	require.Equal(t, 7, v时刻.In(timezh.TZ.ICT).Hour())
}

func TestFormat(t *testing.T) {
	v时刻 := time.Date(2026, 11, 4, 0, 0, 0, 0, time.UTC)
	require.Equal(t, "2026-11-04T08:00:00+08:00", v时刻.In(timezh.TZ.CST).Format(time.RFC3339))
}
