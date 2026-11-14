<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/timezh/release.yml?branch=main&label=BUILD)](https://github.com/yylego/timezh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/timezh)](https://pkg.go.dev/github.com/yylego/timezh)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/timezh/main.svg)](https://coveralls.io/github/yylego/timezh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/timezh.svg)](https://github.com/yylego/timezh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/timezh)](https://goreportcard.com/report/github.com/yylego/timezh)

<!-- TEMPLATE (EN) CLOSE: BADGES -->

# timezh

Comprehensive Go package providing time operations with Chinese function names and intuitive APIs

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

---

## DISCLAIMER

Writing Go code in Chinese is a viable technique, but something to avoid in production engineering. This approach should not be used in serious and business settings. Teams and companies that embrace it could face contempt from peers and negative judgment across the profession. In business companies, this practice is even more prone to becoming a target of public criticism. This project is dedicated to research and academic studies. Do not use this approach in production.

---

## Main Features

- 🎯 **Chinese Function Names**: Chinese-named wrappers around Go's standard time package
- 🔒 **Multiple Time Zones**: Built-in CST, UTC, GMT, and ICT time zones
- 🎨 **Format Options**: Time / date format conversion methods
- 📊 **Round-Trip Conversion**: Time ↔ String ↔ Numeric conversion
- 🌐 **Strict Checks**: The shape must be zero-padded and the date / time must exist, so "2025-13-45" and "9:00:00" get rejected
- 📅 **Date Math**: Add days, list each date in a range, month and week bounds, the ISO 8601 week
- 🧩 **Errors Come Back**: Nothing panics inside; to get a value in one step, wrap the function with `rese.V1`

## Installation

```bash
go get github.com/yylego/timezh
```

## Quick Start

### Time to String Conversion

Convert time objects to various string formats using `TS`.

```go
package main

import (
	"fmt"

	"github.com/yylego/timezh"
)

func main() {
	// Time to String conversion (时间转字符串)
	// Get current date in different formats (获取不同格式的当前日期)
	date := timezh.TS.D日期.Get当日日期(timezh.TZ.CST)
	fmt.Printf("Current Date: %s\n", date) // Output: 2025-05-16

	// Get current time (获取当前时间)
	s时间字符串 := timezh.TS.T时间.Get现在时间(timezh.TZ.CST)
	fmt.Printf("Current Time: %s\n", s时间字符串) // Output: 2025-05-16 19:11:45

	// Get the date before (获取昨天的日期)
	s昨日 := timezh.TS.D日期.Get昨日日期(timezh.TZ.CST)
	fmt.Printf("Date Before: %s\n", s昨日) // Output: 2025-05-15

	// Get 8-digit date format (获取8位数字格式的日期)
	date8 := timezh.TS.D八位数字日期.Get当日日期(timezh.TZ.CST)
	fmt.Printf("8-digit Date: %s\n", date8) // Output: 20250516

	// Get 14-digit time format (获取14位数字格式的时间)
	time14 := timezh.TS.T十四位数字时间.Get现在时间(timezh.TZ.CST)
	fmt.Printf("14-digit Time: %s\n", time14) // Output: 20250516191145
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### String to Time Conversion

Parse time strings into time objects with `Get转为时间` on each `TS` field.

```go
package main

import (
	"fmt"
	"time"

	"github.com/yylego/timezh"
)

func main() {
	// String to Time conversion (字符串转时间)
	// Parse standard time string (解析标准时间字符串)
	t1, err := timezh.TS.T时间.Get转为时间("2025-05-16 19:11:45", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed Time: %v\n", t1)

	// Parse just date (仅解析日期)
	t2, err := timezh.TS.D日期.Get转为时间("2025-05-16", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Date Time: %v\n", t2)

	// Numeric to Time conversion (数字转时间)
	// Create time from date components (从日期组件创建时间)
	t3, err := timezh.NT.New某日零点时间(2025, 5, 16, timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("From Components: %v\n", t3)

	// Timestamp to Time conversion (时间戳转时间)
	// Convert Unix timestamp to time (转换 Unix 时间戳为时间)
	timestamp := time.Now().Unix()
	t4 := timezh.NS.Get秒数时间(timestamp, timezh.TZ.CST)
	fmt.Printf("From Timestamp: %v\n", t4)

	// Time comparison (时间比较)
	// Compare time and time string (比较时间与时间字符串)
	now := time.Now()
	isBefore, err := timezh.CS.Is早于时刻(now, "23:59:59")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Before 23:59:59: %v\n", isBefore)
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

### Time Zone Operations

Work with different time zones and calculate time differences.

```go
package main

import (
	"fmt"

	"github.com/yylego/timezh"
)

func main() {
	// Working with Time Zones (时区操作)
	// Get current date in CST (获取 CST 时区的当前日期)
	dateCST := timezh.TS.D日期.Get当日日期(timezh.TZ.CST)
	fmt.Printf("CST Date: %s\n", dateCST) // Output: 2025-05-16

	// Get current date in UTC (获取 UTC 时区的当前日期)
	dateUTC := timezh.TS.D日期.Get当日日期(timezh.TZ.UTC)
	fmt.Printf("UTC Date: %s\n", dateUTC) // Output: 2025-05-16

	// Get current date in ICT (获取 ICT 时区的当前日期)
	dateICT := timezh.TS.D日期.Get当日日期(timezh.TZ.ICT)
	fmt.Printf("ICT Date: %s\n", dateICT) // Output: 2025-05-16

	// Get recent business days (获取最近的工作日)
	// Get the latest 5 Mon-Fri dates up to now (获取包括今天在内的最近5个工作日)
	workdays := timezh.TS.D日期.Get截至当日工作日期(5, timezh.TZ.CST)
	fmt.Printf("Recent Workdays: %v\n", workdays) // Output: [2025-05-16 2025-05-15 2025-05-14 2025-05-13 2025-05-12]

	// Calculate time difference (计算时间差)
	duration, err := timezh.SD.Get两个时间间隔("2025-05-16 19:00:00", "2025-05-16 20:00:00")
	if err != nil {
		panic(err)
	}
	hours := duration.Hours()
	fmt.Printf("Time Difference: %.0f hours\n", hours) // Output: Time Difference: 1 hours
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

### Working with rese

Functions that can go wrong return an error, and nothing panics inside. When the input is known to be right, [rese](https://github.com/yylego/rese) gives the value in one step, and panics just when the error comes back:

```go
if rese.V1(timezh.CS.Is早于时刻(time.Now(), "23:59:59")) {
	fmt.Println("before 23:59:59")
}
v时间 := rese.V1(timezh.TS.D日期.Get转为时间("2025-05-16", timezh.TZ.CST))
```

## API Reference

### Standard Package First

Since Go 1.20 the `time` package has the layouts `time.DateOnly` ("2006-01-02"), `time.TimeOnly` ("15:04:05") and `time.DateTime` ("2006-01-02 15:04:05"), and the methods of `time.Time` give each part of a date and time. We recommend the standard `time` package first: the Go team maintains it, its docs are the best, and what you learn there works in each project. This package was written to give the reference-time layouts Chinese names that are simple to read, and it stays usable: on top of `time` it adds strict checks (zero-padded shape and a date that exists, so "2025-13-45" and "9:00:00" get rejected), the 8-digit and 6-digit forms, fixed-offset zones, and helpers on Mon-Fri dates.

### Core Modules

#### TS - Time to String (时间转字符串)

Each field formats with one layout:

- `TS.D日期` - "2006-01-02" (`time.DateOnly`)
- `TS.Y年份` - "2006"
- `TS.D月日` - "01-02"
- `TS.T时刻` - "15:04:05" (`time.TimeOnly`)
- `TS.T时分` - "15:04"
- `TS.T时间` - "2006-01-02 15:04:05" (`time.DateTime`)
- `TS.T日期分时` - "2006-01-02 15:04"
- `TS.T国际格式时间` - "2006-01-02T15:04:05" (ISO 8601, the date and time split with the T)
- `TS.T十四位数字时间` - "20060102150405"
- `TS.D八位数字日期` - "20060102"
- `TS.T六位数字时刻` - "150405"

Date fields (`D日期`, `Y年份`, `D月日`, `D八位数字日期`):

- `Get转为文本(t time.Time) string`
- `Get转为时间(s string, z时区位置 *time.Location) (time.Time, error)` - Parse back, the shape must be the same as the layout
- `Match格式(s string) bool` - The shape matches and the date exists
- `Get当日日期(z时区位置 *time.Location) string` - The date of now in the zone
- `Get昨日日期` / `Get明日日期` `(z时区位置 *time.Location) string` - The date before / past the current date, in the zone

The ones below work on the whole date, so use them with `D日期` and `D八位数字日期`; on `Y年份` and `D月日` the results mean nothing:

- `Get加减天数日期(s日期 string, n天数 int) (string, error)` - Move n days, back when n is negative
- `Get区间每天日期(s开始日期, e结束日期 string) ([]string, error)` - Each date from the start to the end, both included
- `Get月初日期` / `Get月末日期` `(t time.Time) string` - Month bounds of t; `Get某月月初日期` / `Get某月月末日期` `(t time.Time, n月数 int) string` move n months first, -1 gives the month before
- `Get当周周一日期` / `Get当周周日日期` `(t time.Time) string` - Week bounds of t, the week runs from Mon to Sun
- `Get截至当日工作日期(n int, z时区位置 *time.Location) []string` - The latest n Mon-Fri dates up to now, newest first
- `Get截至某日工作日期(t time.Time, n int) []string` - The latest n Mon-Fri dates up to t, newest first

Time fields (`T时刻`, `T时分`, `T时间`, `T日期分时`, `T国际格式时间`, `T十四位数字时间`, `T六位数字时刻`):

- `Get转为文本(t time.Time) string`
- `Get转为时间(s string, z时区位置 *time.Location) (time.Time, error)`
- `Match格式(s string) bool`
- `Get现在时间(z时区位置 *time.Location) string`

#### SS - String to String (字符串转字符串)

- `Get拆分日期时刻` - "2006-01-02 15:04:05" → "2006-01-02", "15:04:05"
- `Get标准日期` / `Get八位数字日期` - "20060102" ↔ "2006-01-02"
- `Get取出日期` / `Get取出时刻` - "2006-01-02 15:04:05" → "2006-01-02" / "15:04:05"
- `Get年份` / `Get日期月日` - "2006-01-02" → "2006" / "01-02"
- `Get标准时刻` / `Get六位数字时刻` - "150405" ↔ "15:04:05"
- `Get八位数字转标准时刻` - "15040599" → "15:04:05", the last 2 digits below the second are dropped
- `Get标准时分` - "1504" → "15:04"

#### SI - String to Numbers (字符串转整数)

- `Get年份` / `Get拆分日期` - from "2006-01-02"
- `Get拆分时刻` - from "15:04:05"
- `Get拆分时间` - from "2006-01-02 15:04:05"

#### IS - Numbers to String (整数转字符串)

- `Get日期(年, 月, 日)` → "2006-01-02", `Get时刻(时, 分, 秒)` → "15:04:05", `Get时间(...)` → "2006-01-02 15:04:05", each with an error; numbers that do not make an existing date (such as month 13) give that error, and so do numbers that do not make an existing time

#### NT - Numbers to Time (数字转时间)

- `New某日零点时间(v年份, v月份, v几号 int, z时区位置 *time.Location) (time.Time, error)` - 02-30 gives an error, unlike time.Date, which moves it on to 03-02
- `New当日时刻时间(v几时, v几分, v几秒 int, z时区位置 *time.Location) (time.Time, error)`
- `New当日零点时间(t time.Time) time.Time` - 00:00:00 of the date of t

#### NS - Timestamp to Time (时间戳转时间)

The zone is a must, since the same timestamp gives a different date in each zone:

- `Get秒数时间` / `Get毫秒时间` / `Get微秒时间` / `Get纳秒时间` `(n int64, z时区位置 *time.Location) time.Time` - from a Unix timestamp in each unit

#### TI - Time to Numbers (时间转数量)

- `Get年份` / `Get年份月份` / `Get拆分日期` / `Get拆分时刻` / `Get拆分时间`
- `Get年份周数(t time.Time) (int, int)` - The ISO 8601 week, with its own ISO count of years: 2024-12-30 is week 1 of 2025

#### CS - Compare Time and String (比较时间和字符串)

The strings use the zero-padded "15:04:05" form, so comparing strings gives the same result as comparing times:

- `Is时刻之间(t time.Time, s起始时刻, e结束时刻 string) (bool, error)`
- `Is晚于时刻(t time.Time, s时刻 string) (bool, error)`
- `Is早于时刻(t time.Time, s时刻 string) (bool, error)`

#### SD - String Duration (字符串间隔)

- `Get两个时间间隔(s开始时间, e结束时间 string) (time.Duration, error)` - end minus start
- `Get起点日期间隔` / `Get起点时刻间隔` / `Get起点时间间隔` - from a time to a string
- `Get距今日期间隔` / `Get距今时刻间隔` / `Get距今时间间隔` - from now to a string

#### RS - Regex Match (正则匹配)

Each match checks the zero-padded shape and that the value exists:

- `Match日期` "2006-01-02", `Match时刻` "15:04:05", `Match时间` "2006-01-02 15:04:05", `Match日期分时` "2006-01-02 15:04"
- `Match日期零点` "2006-01-02 00:00:00", `Match八位数字日期` "20060102", `Match六位数字时刻` "150405", `Match八位数字时刻` "15040599", `Match四位数字时分` "1504", `Match年份` "2006"

#### TZ - Time Zones (时区位置)

Fixed-offset zones, which work without the tz database on the machine:

- `TZ.CST` - UTC+8
- `TZ.UTC` - UTC+0 (`time.UTC`)
- `TZ.GMT` - UTC+0
- `TZ.ICT` - UTC+7

#### WK - Week-date (星期)

- `Is工作日(t time.Time) bool` - Mon-Fri
- `Is休息日(t time.Time) bool` - Sat-Sun

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->
