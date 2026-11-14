<!-- TEMPLATE (ZH) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/timezh/release.yml?branch=main&label=BUILD)](https://github.com/yylego/timezh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/timezh)](https://pkg.go.dev/github.com/yylego/timezh)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/timezh/main.svg)](https://coveralls.io/github/yylego/timezh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/timezh.svg)](https://github.com/yylego/timezh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/timezh)](https://goreportcard.com/report/github.com/yylego/timezh)

<!-- TEMPLATE (ZH) CLOSE: BADGES -->

# timezh

全面的 Go 语言时间操作包，使用中文函数名和直观的 API

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) CLOSE: LANGUAGE NAVIGATION -->

---

## 免责声明

使用中文编写 Go 代码在技术上可行，但在实际工程中极不推荐。这种做法不适用于任何正式或商业场景，采用它的团队或公司可能面临同行的轻视和行业的负面评价，在商业公司中尤其可能成为被业内议论的对象。该项目仅供学习和研究，请勿在正式工程中采用。

---

## 主要特性

- 🎯 **中文函数名**: 基于 Go 标准 time 包的中文命名封装
- 🔒 **多时区支持**: 内置 CST、UTC、GMT、ICT 时区
- 🎨 **格式选项**: 时间 / 日期格式转换方法
- 📊 **双向转换**: 时间 ↔ 字符串 ↔ 整数转换
- 🌐 **严格校验**: 必须两位补零、日期和时刻必须真实存在，"2025-13-45"、"9:00:00" 都不通过
- 📅 **日期计算**: 加减天数、列出区间里的每一天、月份和星期的边界、ISO 8601 周数
- 🧩 **返回错误**: 包里不 panic；想一步拿到值，用 `rese.V1` 包一层调用

## 安装

```bash
go get github.com/yylego/timezh
```

## 快速开始

### 时间转字符串

使用 `TS` 将时间对象转换为各种字符串格式。

```go
package main

import (
	"fmt"

	"github.com/yylego/timezh"
)

func main() {
	// 时间转字符串
	// 获取不同格式的当前日期
	date := timezh.TS.D日期.Get当日日期(timezh.TZ.CST)
	fmt.Printf("Current Date: %s\n", date) // 输出: 2025-05-16

	// 获取当前时间
	s时间字符串 := timezh.TS.T时间.Get现在时间(timezh.TZ.CST)
	fmt.Printf("Current Time: %s\n", s时间字符串) // 输出: 2025-05-16 19:11:45

	// 获取昨天的日期
	s昨日 := timezh.TS.D日期.Get昨日日期(timezh.TZ.CST)
	fmt.Printf("Date Before: %s\n", s昨日) // 输出: 2025-05-15

	// 获取8位数字格式的日期
	date8 := timezh.TS.D八位数字日期.Get当日日期(timezh.TZ.CST)
	fmt.Printf("8-digit Date: %s\n", date8) // 输出: 20250516

	// 获取14位数字格式的时间
	time14 := timezh.TS.T十四位数字时间.Get现在时间(timezh.TZ.CST)
	fmt.Printf("14-digit Time: %s\n", time14) // 输出: 20250516191145
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 字符串转时间

用 `TS` 各字段的 `Get转为时间` 把时间字符串解析为时间对象。

```go
package main

import (
	"fmt"
	"time"

	"github.com/yylego/timezh"
)

func main() {
	// 字符串转时间
	// 解析标准时间字符串
	t1, err := timezh.TS.T时间.Get转为时间("2025-05-16 19:11:45", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed Time: %v\n", t1)

	// 仅解析日期
	t2, err := timezh.TS.D日期.Get转为时间("2025-05-16", timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Date Time: %v\n", t2)

	// 数字转时间
	// 从日期组件创建时间
	t3, err := timezh.NT.New某日零点时间(2025, 5, 16, timezh.TZ.CST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("From Components: %v\n", t3)

	// 时间戳转时间
	// 转换 Unix 时间戳为时间
	timestamp := time.Now().Unix()
	t4 := timezh.NS.Get秒数时间(timestamp, timezh.TZ.CST)
	fmt.Printf("From Timestamp: %v\n", t4)

	// 时间比较
	// 比较时间与时间字符串
	now := time.Now()
	isBefore, err := timezh.CS.Is早于时刻(now, "23:59:59")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Before 23:59:59: %v\n", isBefore)
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

### 时区操作

使用不同时区和计算时间差。

```go
package main

import (
	"fmt"

	"github.com/yylego/timezh"
)

func main() {
	// 时区操作
	// 获取 CST 时区的当前日期
	dateCST := timezh.TS.D日期.Get当日日期(timezh.TZ.CST)
	fmt.Printf("CST Date: %s\n", dateCST) // 输出: 2025-05-16

	// 获取 UTC 时区的当前日期
	dateUTC := timezh.TS.D日期.Get当日日期(timezh.TZ.UTC)
	fmt.Printf("UTC Date: %s\n", dateUTC) // 输出: 2025-05-16

	// 获取 ICT 时区的当前日期
	dateICT := timezh.TS.D日期.Get当日日期(timezh.TZ.ICT)
	fmt.Printf("ICT Date: %s\n", dateICT) // 输出: 2025-05-16

	// 获取最近的工作日
	// 获取包括今天在内的最近5个工作日
	workdays := timezh.TS.D日期.Get截至当日工作日期(5, timezh.TZ.CST)
	fmt.Printf("Recent Workdays: %v\n", workdays) // 输出: [2025-05-16 2025-05-15 2025-05-14 2025-05-13 2025-05-12]

	// 计算时间差
	duration, err := timezh.SD.Get两个时间间隔("2025-05-16 19:00:00", "2025-05-16 20:00:00")
	if err != nil {
		panic(err)
	}
	hours := duration.Hours()
	fmt.Printf("Time Difference: %.0f hours\n", hours) // 输出: Time Difference: 1 hours
}
```

⬆️ **源码:** [源码](internal/demos/demo3x/main.go)

### 搭配 rese 使用

可能出错的函数都返回 error，包里不 panic。确定输入没问题时，用 [rese](https://github.com/yylego/rese) 一步拿到值，只有真出错时才 panic：

```go
if rese.V1(timezh.CS.Is早于时刻(time.Now(), "23:59:59")) {
	fmt.Println("before 23:59:59")
}
v时间 := rese.V1(timezh.TS.D日期.Get转为时间("2025-05-16", timezh.TZ.CST))
```

## API 说明

### 标准库优先

Go 1.20 起标准库有了 `time.DateOnly`（"2006-01-02"）、`time.TimeOnly`（"15:04:05"）、`time.DateTime`（"2006-01-02 15:04:05"）这几个布局，取年月日、时分秒用 `time.Time` 自带的方法就行。我们推荐优先用标准库：它由 Go 官方维护、文档最全，学会了在每个项目里都用得上。本包写于刚接触 Go、觉得参考时间布局难记的时候，用中文名字把常用格式起了名，现在照常可用：在标准库之上，它还提供严格校验（必须两位补零、必须是真实存在的日期，"2025-13-45"、"9:00:00" 都不通过）、8 位和 6 位数字格式、固定偏移的时区和工作日工具。

### 核心模块

#### TS - 时间转字符串

每个字段按一种布局格式化：

- `TS.D日期` - "2006-01-02"（`time.DateOnly`）
- `TS.Y年份` - "2006"
- `TS.D月日` - "01-02"
- `TS.T时刻` - "15:04:05"（`time.TimeOnly`）
- `TS.T时分` - "15:04"
- `TS.T时间` - "2006-01-02 15:04:05"（`time.DateTime`）
- `TS.T日期分时` - "2006-01-02 15:04"
- `TS.T国际格式时间` - "2006-01-02T15:04:05"（ISO 8601，日期和时刻用 T 隔开）
- `TS.T十四位数字时间` - "20060102150405"
- `TS.D八位数字日期` - "20060102"
- `TS.T六位数字时刻` - "150405"

日期字段（`D日期`、`Y年份`、`D月日`、`D八位数字日期`）：

- `Get转为文本(t time.Time) string`
- `Get转为时间(s string, z时区位置 *time.Location) (time.Time, error)` - 解析回时间，格式必须完全一致
- `Match格式(s string) bool` - 格式一致、日期真实存在
- `Get当日日期(z时区位置 *time.Location) string` - 该时区的今天
- `Get昨日日期` / `Get明日日期` `(z时区位置 *time.Location) string` - 该时区的昨天 / 明天

下面这些按完整日期计算，要用在 `D日期`、`D八位数字日期` 上；用在 `Y年份`、`D月日` 上得到的结果没有意义：

- `Get加减天数日期(s日期 string, n天数 int) (string, error)` - 往后挪 n 天，n 为负数时往前挪
- `Get区间每天日期(s开始日期, e结束日期 string) ([]string, error)` - 从开始到结束的每一天，两头都包含
- `Get月初日期` / `Get月末日期` `(t time.Time) string` - t 所在月份的边界；`Get某月月初日期` / `Get某月月末日期` `(t time.Time, n月数 int) string` 先往后挪 n 个月，-1 就是上个月
- `Get当周周一日期` / `Get当周周日日期` `(t time.Time) string` - t 所在那一周的边界，一周从周一算到周日
- `Get截至当日工作日期(n int, z时区位置 *time.Location) []string` - 到今天为止最近 n 个工作日（周一到周五），新的在前
- `Get截至某日工作日期(t time.Time, n int) []string` - 到 t 为止最近 n 个工作日，新的在前

时间字段（`T时刻`、`T时分`、`T时间`、`T日期分时`、`T国际格式时间`、`T十四位数字时间`、`T六位数字时刻`）：

- `Get转为文本(t time.Time) string`
- `Get转为时间(s string, z时区位置 *time.Location) (time.Time, error)`
- `Match格式(s string) bool`
- `Get现在时间(z时区位置 *time.Location) string`

#### SS - 字符串转字符串

- `Get拆分日期时刻` - "2006-01-02 15:04:05" → "2006-01-02"、"15:04:05"
- `Get标准日期` / `Get八位数字日期` - "20060102" ↔ "2006-01-02"
- `Get取出日期` / `Get取出时刻` - "2006-01-02 15:04:05" → "2006-01-02" / "15:04:05"
- `Get年份` / `Get日期月日` - "2006-01-02" → "2006" / "01-02"
- `Get标准时刻` / `Get六位数字时刻` - "150405" ↔ "15:04:05"
- `Get八位数字转标准时刻` - "15040599" → "15:04:05"，末尾两位秒以下的小数去掉
- `Get标准时分` - "1504" → "15:04"

#### SI - 字符串转整数

- `Get年份` / `Get拆分日期` - 从 "2006-01-02" 取
- `Get拆分时刻` - 从 "15:04:05" 取
- `Get拆分时间` - 从 "2006-01-02 15:04:05" 取

#### IS - 整数转字符串

- `Get日期(年, 月, 日)` → "2006-01-02"，`Get时刻(时, 分, 秒)` → "15:04:05"，`Get时间(...)` → "2006-01-02 15:04:05"，都带一个 error；拼不成真实日期或时刻的数字（如 13 月）返回错误

#### NT - 数字转时间

- `New某日零点时间(v年份, v月份, v几号 int, z时区位置 *time.Location) (time.Time, error)` - 2 月 30 日返回错误，不像 time.Date 那样顺延成 3 月 2 日
- `New当日时刻时间(v几时, v几分, v几秒 int, z时区位置 *time.Location) (time.Time, error)`
- `New当日零点时间(t time.Time) time.Time` - t 那一天的 00:00:00

#### NS - 时间戳转时间

必须传时区，因为同一个时间戳在不同时区里日期可能不同：

- `Get秒数时间` / `Get毫秒时间` / `Get微秒时间` / `Get纳秒时间` `(n int64, z时区位置 *time.Location) time.Time` - 从对应单位的 Unix 时间戳转换

#### TI - 时间转数量

- `Get年份` / `Get年份月份` / `Get拆分日期` / `Get拆分时刻` / `Get拆分时间`
- `Get年份周数(t time.Time) (int, int)` - ISO 8601 周数，连同它所属的年份一起返回：2024-12-30 是 2025 年第 1 周

#### CS - 比较时间和字符串

字符串用两位补零的 "15:04:05" 格式，这样字符串的大小顺序就是时间的先后顺序：

- `Is时刻之间(t time.Time, s起始时刻, e结束时刻 string) (bool, error)`
- `Is晚于时刻(t time.Time, s时刻 string) (bool, error)`
- `Is早于时刻(t time.Time, s时刻 string) (bool, error)`

#### SD - 字符串间隔

- `Get两个时间间隔(s开始时间, e结束时间 string) (time.Duration, error)` - 结束减开始
- `Get起点日期间隔` / `Get起点时刻间隔` / `Get起点时间间隔` - 从某个时间到字符串表示的时间
- `Get距今日期间隔` / `Get距今时刻间隔` / `Get距今时间间隔` - 从现在到字符串表示的时间

#### RS - 正则匹配

每个匹配都查两位补零的格式，也查取值是否真实存在：

- `Match日期` "2006-01-02"，`Match时刻` "15:04:05"，`Match时间` "2006-01-02 15:04:05"，`Match日期分时` "2006-01-02 15:04"
- `Match日期零点` "2006-01-02 00:00:00"，`Match八位数字日期` "20060102"，`Match六位数字时刻` "150405"，`Match八位数字时刻` "15040599"，`Match四位数字时分` "1504"，`Match年份` "2006"

#### TZ - 时区位置

固定偏移的时区，机器上没有时区数据库也能用：

- `TZ.CST` - UTC+8
- `TZ.UTC` - UTC+0（`time.UTC`）
- `TZ.GMT` - UTC+0
- `TZ.ICT` - UTC+7

#### WK - 星期

- `Is工作日(t time.Time) bool` - 周一到周五
- `Is休息日(t time.Time) bool` - 周六、周日

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) CLOSE: STANDARD PROJECT FOOTER -->
