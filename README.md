[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/timezh/release.yml?branch=main&label=BUILD)](https://github.com/yylego/timezh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/timezh)](https://pkg.go.dev/github.com/yylego/timezh)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/timezh/main.svg)](https://coveralls.io/github/yylego/timezh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://github.com/yylego/timezh)
[![GitHub Release](https://img.shields.io/github/release/yylego/timezh.svg)](https://github.com/yylego/timezh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/timezh)](https://goreportcard.com/report/github.com/yylego/timezh)

# timezh

Comprehensive Go package providing time operations with Chinese function names and intuitive APIs

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

---

## DISCLAIMER

Writing Go code in Chinese is a viable technique, but something to avoid in production engineering. This approach should not be used in serious and business settings. Teams and companies that embrace it could face contempt from peers and negative judgment across the profession. In business companies, this practice is even more prone to becoming a target of public criticism. This project is dedicated to research and academic studies. Do not use this approach in production.

---

## Main Features

🎯 **Chinese Function Names**: Intuitive Chinese-named wrappers around Go's standard time package
🔒 **Multiple Time Zones**: Built-in support on CST, UTC, GMT, and ICT time zones
🎨 **Rich Format Options**: Various time/date format conversion methods
📊 **Bidirectional Conversion**: Time ↔ String ↔ Numeric seamless conversion
🌐 **Regex Matching**: Smart time string pattern recognition and extraction

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
	date := timezh.TS.D日期.Get现在日期(timezh.TZ.CST)
	fmt.Printf("Current Date: %s\n", date) // Output: 2025-05-16

	// Get current time (获取当前时间)
	s时间字符串 := timezh.TS.T时间.Get现在时间(timezh.TZ.CST)
	fmt.Printf("Current Time: %s\n", s时间字符串) // Output: 2025-05-16 19:11:45

	// Get yesterday's date (获取昨天的日期)
	yesterday := timezh.TS.D日期.Get昨日日期(timezh.TZ.CST)
	fmt.Printf("Yesterday: %s\n", yesterday) // Output: 2025-05-15

	// Get 8-digit date format (获取8位数字格式的日期)
	date8 := timezh.TS.D日期8位数.Get现在日期(timezh.TZ.CST)
	fmt.Printf("8-digit Date: %s\n", date8) // Output: 20250516

	// Get 14-digit time format (获取14位数字格式的时间)
	time14 := timezh.TS.T时间14位数.Get现在时间(timezh.TZ.CST)
	fmt.Printf("14-digit Time: %s\n", time14) // Output: 20250516191145
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### String to Time Conversion

Parse time strings into time objects using `ST`.

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
	t1, err := timezh.ST.Get时间("2025-05-16 19:11:45", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed Time: %v\n", t1)

	// Parse just date (仅解析日期)
	t2, err := timezh.ST.Get通过日期得到时间("2025-05-16", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Date Time: %v\n", t2)

	// Numeric to Time conversion (数字转时间)
	// Create time from date components (从日期组件创建时间)
	t3 := timezh.NT.New由年月日得到零时时间(2025, 5, 16, timezh.TZ.CST)
	fmt.Printf("From Components: %v\n", t3)

	// Timestamp to Time conversion (时间戳转时间)
	// Convert Unix timestamp to time (转换 Unix 时间戳为时间)
	timestamp := time.Now().Unix()
	t4 := timezh.NS.Get秒数时间(timestamp)
	fmt.Printf("From Timestamp: %v\n", t4)

	// Time comparison (时间比较)
	// Compare time and time string (比较时间与时间字符串)
	now := time.Now()
	isBefore := timezh.CS.B时间在时分秒后面(now, "23:59:59")
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
	dateCST := timezh.TS.D日期.Get现在日期(timezh.TZ.CST)
	fmt.Printf("CST Date: %s\n", dateCST) // Output: 2025-05-16

	// Get current date in UTC (获取 UTC 时区的当前日期)
	dateUTC := timezh.TS.D日期.Get现在日期(timezh.TZ.UTC)
	fmt.Printf("UTC Date: %s\n", dateUTC) // Output: 2025-05-16

	// Get current date in ICT (获取 ICT 时区的当前日期)
	dateICT := timezh.TS.D日期.Get现在日期(timezh.TZ.ICT)
	fmt.Printf("ICT Date: %s\n", dateICT) // Output: 2025-05-16

	// Get recent business days (获取最近的工作日)
	// Get the most recent 5 business days including today (获取包括今天在内的最近5个工作日)
	workdays := timezh.TS.D日期.Get最近含当日的前N个工作日(5, timezh.TZ.CST)
	fmt.Printf("Recent Workdays: %v\n", workdays) // Output: ["2025-05-16", "2025-05-15", "2025-05-14", "2025-05-13", "2025-05-12"]

	// Calculate time difference (计算时间差)
	duration, err := timezh.SD.Get两字符串时间间隔("2025-05-16 20:00:00", "2025-05-16 19:00:00")
	if err != nil {
		panic(err)
	}
	hours := duration.Hours()
	fmt.Printf("Time Difference: %.0f hours\n", hours) // Output: 1 hours
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

## API Reference

### Core Modules

#### TS - Time to String (时间转字符串)

Time formatting module with predefined format templates:

- `TS.D日期` - Date format "2006-01-02"
- `TS.Y年份` - Just the "2006" component
- `TS.D月日` - Month-Day "01-02"
- `TS.T时刻` - Just time "15:04:05"
- `TS.T时间` - Complete datetime "2006-01-02 15:04:05"
- `TS.T日期和T和时刻` - ISO 8601 format "2006-01-02T15:04:05"
- `TS.T时间14位数` - 14-digit format "20060102150405"
- `TS.D日期8位数` - 8-digit date "20060102"
- `TS.T时刻6位数` - 6-digit time "150405"

**Main Methods:**
- `Get转字符串(t time.Time) string` - Convert time to string
- `Get现在日期(z时区位置 *time.Location) string` - Get current date
- `Get昨日日期(z时区位置 *time.Location) string` - Get yesterday's date
- `Get最近含当日的前N个工作日(n int, z时区位置 *time.Location) []string` - Get recent N business days

#### ST - String to Time (字符串转时间)

String parsing module:

- `Get时间(s标准时间 string, z时区位置 *time.Location) (time.Time, error)` - Parse "2006-01-02 15:04:05"
- `Get日期和分时(s标准时间没有秒数 string, z时区位置 *time.Location) (time.Time, error)` - Parse "2006-01-02 15:04"
- `Get通过日期得到时间(s标准日期 string, z时区位置 *time.Location) (time.Time, error)` - Parse "2006-01-02"
- `Get通过时刻得到当日时间(s标准时刻 string, z时区位置 *time.Location) (time.Time, error)` - Parse "15:04:05"

#### NT - Numeric to Time (数字转时间)

Construct time from numeric components:

- `New由年月日得到零时时间(year, month, day int, z时区位置 *time.Location) time.Time`
- `New由当日时分秒得到时间(hour, minute, second int, z时区位置 *time.Location) time.Time`

#### NS - Timestamp to Time (时间戳转时间)

Convert Unix timestamps to time objects:

- `Get秒数时间(sn总秒数 int64) time.Time`
- `Get毫秒时间(ms微秒数 int64) time.Time`
- `Get微秒时间(ms微秒数 int64) time.Time`
- `Get纳秒时间(ns纳秒数 int64) time.Time`

#### TI - Time to Numeric (时间转数量)

Extract numeric components from time:

- `Get年份(t time.Time) int`
- `Get年份和月份(t time.Time) (int, int)`
- `Get年份月份几号(t time.Time) (int, int, int)`
- `Get几时几分秒数(t time.Time) (int, int, int)`
- `Get数量(t time.Time) (int, int, int, int, int, int)` - Extract components

#### CS - Compare Strings (比较字符串)

Compare time strings:

- `I在时分秒内含边界(t time.Time, s起始时刻, e结束时刻 string) bool`
- `A时间在时分秒前面(t time.Time, s时刻 string) bool`
- `B时间在时分秒后面(t time.Time, s时刻 string) bool`

#### SD - String Duration (字符串间隔)

Calculate time differences between strings:

- `Get两字符串时间间隔(s开始时间, e结束时间 string) (time.Duration, error)`

#### RS - Regex String (正则匹配)

Extract time patterns from strings:

- `Get匹配时间(s string) string` - Match "2006-01-02 15:04:05"
- `Get匹配日期(s string) string` - Match "2006-01-02"
- `Get匹配时刻(s string) string` - Match "15:04:05"

#### TZ - Time Zones (时区位置)

Built-in time zone locations:

- `TZ.CST` - China Standard Time (北京时间) UTC+8
- `TZ.UTC` - Coordinated Universal Time (世界标准时间) UTC+0
- `TZ.GMT` - Greenwich Mean Time (格林尼治标准时间) UTC+0
- `TZ.ICT` - Indochina Time (东南亚时间) UTC+7

#### WK - Week-date (星期)

Week-date operations:

- `Is工作日(t time.Time) bool` - Check if business day (Mon-Fri)
- `Is周六日(t time.Time) bool` - Check if weekend (Sat-Sun)

## Examples

### Working with Time Zones

```go
// CST - China Standard Time
dateCST := timezh.TS.D日期.Get现在日期(timezh.TZ.CST)

// UTC - Universal Coordinated Time
dateUTC := timezh.TS.D日期.Get现在日期(timezh.TZ.UTC)

// ICT - Indochina Time
dateICT := timezh.TS.D日期.Get现在日期(timezh.TZ.ICT)
```

### Get Recent Business Days

```go
// Get the most recent 5 business days including today
workdays := timezh.TS.D日期.Get最近含当日的前N个工作日(5, timezh.TZ.CST)
// Output: ["2025-05-16", "2025-05-15", "2025-05-14", "2025-05-13", "2025-05-12"]
```

### Extract Time from String

```go
text := "Meeting scheduled on 2025-05-16 at 14:30:00"
s时间字符串 := timezh.RS.Get匹配时间(text)
// Output: "2025-05-16 14:30:00"
```

### Calculate Time Difference

```go
duration, err := timezh.SD.Get两字符串时间间隔("2025-05-16 20:00:00", "2025-05-16 19:00:00")
if err != nil {
	panic(err)
}
hours := duration.Hours()
// Output: 1 h
```

## Design Concept

This package follows these principles:

1. **Chinese Naming**: Functions use intuitive Chinese names matching usage scenarios
2. **Type-Safe Operations**: Leverages Go's type system to achieve safe time operations
3. **Time Zone Awareness**: Explicit time zone handling across operations
4. **Flexible Formats**: Multiple pre-defined and custom format support
5. **Bidirectional Conversion**: Seamless conversion between time, string, and numeric formats

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
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
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
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

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/timezh.svg?variant=adaptive)](https://starchart.cc/yylego/timezh)
