package timezh_test

import (
	"testing"
	"time"

	"github.com/yylego/timezh"
)

func Test_cs比较时间和字符串_I在时分秒内含边界(t *testing.T) {
	t.Log(timezh.CS.I在时分秒内含边界(time.Now().In(timezh.TZ.CST), "09:15:00", "15:00:00"))
}
