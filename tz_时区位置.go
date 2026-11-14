package timezh

import "time"

// TZ provides often used fixed-offset time zones: CST (UTC+8), UTC, GMT (UTC+0) and ICT (UTC+7)
//
// TZ 提供常用的固定偏移时区：CST（UTC+8）、UTC、GMT（UTC+0）和 ICT（UTC+7）
var TZ = struct {
	CST *time.Location // UTC+8 | UTC+8
	UTC *time.Location // UTC+0 | 协调世界时，UTC+0
	GMT *time.Location // Greenwich Mean Time, UTC+0 | 格林尼治标准时间，UTC+0
	ICT *time.Location // UTC+7 | UTC+7
}{
	CST: time.FixedZone("CST", 8*3600),
	UTC: time.UTC,
	GMT: time.FixedZone("GMT", 0),
	ICT: time.FixedZone("ICT", 7*3600),
}
