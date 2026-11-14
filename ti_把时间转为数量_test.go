package timezh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yylego/timezh"
)

func Test_ti把时间转为数量_Get拆分日期(t *testing.T) {
	v年份, v月份, v几号 := timezh.TI.Get拆分日期(time.Date(2026, 11, 4, 15, 4, 5, 0, timezh.TZ.CST))
	require.Equal(t, []int{2026, 11, 4}, []int{v年份, v月份, v几号})
}

func Test_ti把时间转为数量_Get拆分时刻(t *testing.T) {
	v几时, v几分, v秒数 := timezh.TI.Get拆分时刻(time.Date(2026, 11, 4, 15, 4, 5, 0, timezh.TZ.CST))
	require.Equal(t, []int{15, 4, 5}, []int{v几时, v几分, v秒数})
}

func Test_ti把时间转为数量_Get拆分时间(t *testing.T) {
	v年份, v月份, v几号, v几时, v几分, v秒数 := timezh.TI.Get拆分时间(time.Date(2026, 11, 4, 15, 4, 5, 0, timezh.TZ.CST))
	require.Equal(t, []int{2026, 11, 4, 15, 4, 5}, []int{v年份, v月份, v几号, v几时, v几分, v秒数})
}

func Test_ti把时间转为数量_Get年份月份(t *testing.T) {
	v年份, v月份 := timezh.TI.Get年份月份(time.Date(2026, 11, 4, 15, 4, 5, 0, timezh.TZ.CST))
	require.Equal(t, []int{2026, 11}, []int{v年份, v月份})
}

func Test_ti把时间转为数量_Get年份周数(t *testing.T) {
	v年份, v周数 := timezh.TI.Get年份周数(time.Date(2026, 11, 4, 15, 4, 5, 0, timezh.TZ.CST))
	require.Equal(t, []int{2026, 45}, []int{v年份, v周数})

	// Around Jan 1 the week can belong to the next ISO count of years, and to the one before
	// 年头年尾，周数可能算到下一年或上一年
	v年份, v周数 = timezh.TI.Get年份周数(time.Date(2024, 12, 30, 9, 0, 0, 0, timezh.TZ.CST))
	require.Equal(t, []int{2025, 1}, []int{v年份, v周数})
	v年份, v周数 = timezh.TI.Get年份周数(time.Date(2027, 1, 1, 9, 0, 0, 0, timezh.TZ.CST))
	require.Equal(t, []int{2026, 53}, []int{v年份, v周数})
}

func Test_ti把时间转为数量_Get年份(t *testing.T) {
	require.Equal(t, 2026, timezh.TI.Get年份(time.Date(2026, 11, 4, 15, 4, 5, 0, timezh.TZ.CST)))
}
