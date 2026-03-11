package timezh

import "time"

// TZ provides often used timezone locations
// Includes CST (China), UTC (Coordinated), GMT (Greenwich), and ICT (Indochina) time zones
//
// TZ 提供常用时区位置
// 包括 CST（中国）、UTC（世界标准）、GMT（格林尼治）和 ICT（中南半岛）时区
var TZ = struct {
	CST *time.Location // China Standard Time (Beijing Time) UTC+8 | 中国标准时间（北京时间）
	UTC *time.Location // Coordinated Universal Time UTC+0 | 协调世界时
	GMT *time.Location // Greenwich Mean Time UTC+0 | 格林尼治标准时间
	ICT *time.Location // Indochina Time (Thailand, Vietnam, Cambodia) UTC+7 | 中南半岛时间（泰国、越南、柬埔寨等）
}{
	CST: time.FixedZone("CST", 8*3600), // China Standard Time UTC+8 | 中国标准时间
	UTC: time.UTC,                      // Coordinated Universal Time UTC+0 | 协调世界时
	GMT: time.FixedZone("GMT", 0),      // Greenwich Mean Time UTC+0 | 格林尼治标准时间
	ICT: time.FixedZone("ICT", 7*3600), // Indochina Time UTC+7 | 中南半岛时间
}
