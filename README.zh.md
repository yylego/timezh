[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/timezh/release.yml?branch=main&label=BUILD)](https://github.com/yylego/timezh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/timezh)](https://pkg.go.dev/github.com/yylego/timezh)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/timezh/main.svg)](https://coveralls.io/github/yylego/timezh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://github.com/yylego/timezh)
[![GitHub Release](https://img.shields.io/github/release/yylego/timezh.svg)](https://github.com/yylego/timezh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/timezh)](https://goreportcard.com/report/github.com/yylego/timezh)

# timezh

全面的 Go 语言时间操作包，使用中文函数名和直观的 API

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

---

## 免责声明

使用中文编写 Go 代码在技术上可行，但在实际工程中极不推荐。这种做法不适用于任何正式或商业场景，采用它的团队或公司可能面临同行的轻视和行业的负面评价，在商业公司中尤其可能成为被业内议论的对象。该项目仅供学习和研究，请勿在正式工程中采用。

---

## 主要特性

🎯 **中文函数名**: 基于 Go 标准 time 包的直观中文命名封装
🔒 **多时区支持**: 内置 CST、UTC、GMT、ICT 时区支持
🎨 **丰富格式选项**: 多种时间/日期格式转换方法
📊 **双向转换**: 时间 ↔ 字符串 ↔ 整数无缝转换
🌐 **正则匹配**: 智能时间字符串模式识别和提取

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
	date := timezh.TS.D日期.Get现在日期(timezh.TZ.CST)
	fmt.Printf("当前日期: %s\n", date) // 输出: 2025-05-16

	// 获取当前时间
	s时间字符串 := timezh.TS.T时间.Get现在时间(timezh.TZ.CST)
	fmt.Printf("当前时间: %s\n", s时间字符串) // 输出: 2025-05-16 19:11:45

	// 获取昨天的日期
	yesterday := timezh.TS.D日期.Get昨日日期(timezh.TZ.CST)
	fmt.Printf("昨天: %s\n", yesterday) // 输出: 2025-05-15

	// 获取8位数字格式的日期
	date8 := timezh.TS.D日期8位数.Get现在日期(timezh.TZ.CST)
	fmt.Printf("8位数日期: %s\n", date8) // 输出: 20250516

	// 获取14位数字格式的时间
	time14 := timezh.TS.T时间14位数.Get现在时间(timezh.TZ.CST)
	fmt.Printf("14位数时间: %s\n", time14) // 输出: 20250516191145
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 字符串转时间

使用 `ST` 将时间字符串解析为时间对象。

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
	t1 := timezh.ST.Get时间("2025-05-16 19:11:45", timezh.TZ.CST)
	fmt.Printf("解析的时间: %v\n", t1)

	// 仅解析日期
	t2 := timezh.ST.Get通过日期得到时间("2025-05-16", timezh.TZ.CST)
	fmt.Printf("日期时间: %v\n", t2)

	// 数字转时间
	// 从日期组件创建时间
	t3 := timezh.NT.New由年月日得到零时时间(2025, 5, 16, timezh.TZ.CST)
	fmt.Printf("从组件创建: %v\n", t3)

	// 时间戳转时间
	// 转换 Unix 时间戳为时间
	timestamp := time.Now().Unix()
	t4 := timezh.NS.Get秒数时间(timestamp)
	fmt.Printf("从时间戳: %v\n", t4)

	// 时间比较
	// 比较时间与时间字符串
	now := time.Now()
	isBefore := timezh.CS.B时间在时分秒后面(now, "23:59:59")
	fmt.Printf("在 23:59:59 之前: %v\n", isBefore)
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
	dateCST := timezh.TS.D日期.Get现在日期(timezh.TZ.CST)
	fmt.Printf("CST 日期: %s\n", dateCST) // 输出: 2025-05-16

	// 获取 UTC 时区的当前日期
	dateUTC := timezh.TS.D日期.Get现在日期(timezh.TZ.UTC)
	fmt.Printf("UTC 日期: %s\n", dateUTC) // 输出: 2025-05-16

	// 获取 ICT 时区的当前日期
	dateICT := timezh.TS.D日期.Get现在日期(timezh.TZ.ICT)
	fmt.Printf("ICT 日期: %s\n", dateICT) // 输出: 2025-05-16

	// 获取最近的工作日
	// 获取包括今天在内的最近5个工作日
	workdays := timezh.TS.D日期.Get最近含当日的前N个工作日(5, timezh.TZ.CST)
	fmt.Printf("最近工作日: %v\n", workdays) // 输出: ["2025-05-16", "2025-05-15", "2025-05-14", "2025-05-13", "2025-05-12"]

	// 计算时间差
	duration := timezh.SD.Get两字符串时间间隔("2025-05-16 20:00:00", "2025-05-16 19:00:00")
	hours := duration.Hours()
	fmt.Printf("时间差: %.0f 小时\n", hours) // 输出: 1 小时
}
```

⬆️ **源码:** [源码](internal/demos/demo3x/main.go)

## API 参考

### 核心模块

#### TS - 时间转字符串

时间格式化模块，预定义格式模板：

- `TS.D日期` - 日期格式 "2006-01-02"
- `TS.Y年份` - 仅年份 "2006"
- `TS.D月日` - 月-日 "01-02"
- `TS.T时刻` - 仅时间 "15:04:05"
- `TS.T时间` - 完整日期时间 "2006-01-02 15:04:05"
- `TS.T日期和T和时刻` - ISO 8601 格式 "2006-01-02T15:04:05"
- `TS.T时间14位数` - 14位数字格式 "20060102150405"
- `TS.D日期8位数` - 8位数日期 "20060102"
- `TS.T时刻6位数` - 6位数时间 "150405"

**主要方法:**
- `Get转字符串(t time.Time) string` - 转换时间为字符串
- `Get现在日期(z时区位置 *time.Location) string` - 获取当前日期
- `Get昨日日期(z时区位置 *time.Location) string` - 获取昨天日期
- `Get最近含当日的前N个工作日(n int, z时区位置 *time.Location) []string` - 获取最近 N 个工作日

#### ST - 字符串转时间

字符串解析模块：

- `Get时间(s标准时间 string, z时区位置 *time.Location) time.Time` - 解析 "2006-01-02 15:04:05"
- `Get日期和分时(s标准时间没有秒数 string, z时区位置 *time.Location) time.Time` - 解析 "2006-01-02 15:04"
- `Get通过日期得到时间(s标准日期 string, z时区位置 *time.Location) time.Time` - 解析 "2006-01-02"
- `Get通过时刻得到当日时间(s标准时刻 string, z时区位置 *time.Location) time.Time` - 解析 "15:04:05"

#### NT - 数字转时间

从数字组件构建时间：

- `New由年月日得到零时时间(year, month, day int, z时区位置 *time.Location) time.Time`
- `New由当日时分秒得到时间(hour, minute, second int, z时区位置 *time.Location) time.Time`

#### NS - 时间戳转时间

转换 Unix 时间戳为时间对象：

- `Get秒数时间(sn总秒数 int64) time.Time`
- `Get毫秒时间(ms微秒数 int64) time.Time`
- `Get微秒时间(ms微秒数 int64) time.Time`
- `Get纳秒时间(ns纳秒数 int64) time.Time`

#### TI - 时间转数量

从时间提取数字组件：

- `Get年份(t time.Time) int`
- `Get年份和月份(t time.Time) (int, int)`
- `Get年份月份几号(t time.Time) (int, int, int)`
- `Get几时几分秒数(t time.Time) (int, int, int)`
- `Get数量(t time.Time) (int, int, int, int, int, int)` - 提取所有组件

#### CS - 比较字符串

比较时间字符串：

- `I在时分秒内含边界(t time.Time, s起始时刻, e结束时刻 string) bool`
- `A时间在时分秒前面(t time.Time, s时刻 string) bool`
- `B时间在时分秒后面(t time.Time, s时刻 string) bool`

#### SD - 字符串间隔

计算字符串间的时间差：

- `Get计算时间差(s1, s2 string, z时区位置 *time.Location) time.Duration`

#### RS - 正则匹配

从字符串中提取时间模式：

- `Get匹配时间(s string) string` - 匹配 "2006-01-02 15:04:05"
- `Get匹配日期(s string) string` - 匹配 "2006-01-02"
- `Get匹配时刻(s string) string` - 匹配 "15:04:05"

#### TZ - 时区位置

内置时区位置：

- `TZ.CST` - 中国标准时间（北京时间）UTC+8
- `TZ.UTC` - 协调世界时（世界标准时间）UTC+0
- `TZ.GMT` - 格林尼治标准时间 UTC+0
- `TZ.ICT` - 中南半岛时间（东南亚时间）UTC+7

#### WK - 星期

星期操作：

- `Is工作日(t time.Time) bool` - 检查是否工作日（周一至周五）
- `Is周六日(t time.Time) bool` - 检查是否周末（周六和周日）

## 示例

### 使用时区

```go
// CST - 中国标准时间
dateCST := timezh.TS.D日期.Get现在日期(timezh.TZ.CST)

// UTC - 协调世界时
dateUTC := timezh.TS.D日期.Get现在日期(timezh.TZ.UTC)

// ICT - 中南半岛时间
dateICT := timezh.TS.D日期.Get现在日期(timezh.TZ.ICT)
```

### 获取最近工作日

```go
// 获取包括今天在内的最近5个工作日
workdays := timezh.TS.D日期.Get最近含当日的前N个工作日(5, timezh.TZ.CST)
// 输出: ["2025-05-16", "2025-05-15", "2025-05-14", "2025-05-13", "2025-05-12"]
```

### 从字符串提取时间

```go
text := "会议安排在 2025-05-16 at 14:30:00"
s时间字符串 := timezh.RS.Get匹配时间(text)
// 输出: "2025-05-16 14:30:00"
```

### 计算时间差

```go
duration := timezh.SD.Get计算时间差("2025-05-16 20:00:00", "2025-05-16 19:00:00", timezh.TZ.CST)
hours := duration.Hours()
// 输出: 1 小时
```

## 设计理念

本包遵循以下原则：

1. **中文命名**: 所有函数使用符合使用场景的直观中文名称
2. **类型安全**: 利用 Go 的类型系统进行安全的时间操作
3. **时区感知**: 所有操作中显式处理时区
4. **格式灵活**: 支持多种预定义和自定义格式
5. **双向转换**: 时间、字符串和数字格式间的无缝转换

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
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
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yylego/timezh.svg?variant=adaptive)](https://starchart.cc/yylego/timezh)
