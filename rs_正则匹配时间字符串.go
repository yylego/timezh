package timezh

import (
	"regexp"
	"strings"
	"time"

	"github.com/yylego/erero"
)

// RS is the regex-based time string component
// Each match checks the exact shape (fixed digits, zero-padded) and then that the value exists as a date/time,
// so "2025-13-45" and "25:61:61" do not match
//
// RS 是全局基于正则的时间字符串匹配器
// 每个匹配先查格式（位数固定、两位补零），再查值是不是真实存在的日期/时刻，
// 所以 "2025-13-45"、"25:61:61" 都不匹配
var RS = &rs正则匹配时间字符串{
	v日期:     newRs布局格式(time.DateOnly),
	v时刻:     newRs布局格式(time.TimeOnly),
	v时间:     newRs布局格式(time.DateTime),
	v日期分时:   newRs布局格式("2006-01-02 15:04"),
	v日期零点:   newRs时间格式(`^\d{4}-\d{2}-\d{2} 00:00:00$`, time.DateTime),
	v八位数字日期: newRs布局格式("20060102"),
	v六位数字时刻: newRs布局格式("150405"),
	v四位数字时分: newRs布局格式("1504"),
	v年份:     newRs布局格式("2006"),
	v8位数字:   regexp.MustCompile(`^\d{8}$`),
}

// rs正则匹配时间字符串 validates time strings using regex patterns plus a parse of the value
//
// rs正则匹配时间字符串 用正则检查格式，再解析一遍检查取值
type rs正则匹配时间字符串 struct {
	v日期     *rs时间格式 // "2006-01-02"
	v时刻     *rs时间格式 // "15:04:05"
	v时间     *rs时间格式 // "2006-01-02 15:04:05"
	v日期分时   *rs时间格式 // "2006-01-02 15:04"
	v日期零点   *rs时间格式 // "2006-01-02 00:00:00"
	v八位数字日期 *rs时间格式 // "20060102"
	v六位数字时刻 *rs时间格式 // "150405"
	v四位数字时分 *rs时间格式 // "1504"
	v年份     *rs时间格式 // "2006"
	v8位数字   *regexp.Regexp
}

// rs时间格式 pairs the regex that fixes the shape with the Go layout that checks the value
// The regex is needed since time.Parse accepts "9:04:05" with the layout "15:04:05"
//
// rs时间格式 把固定格式的正则和检查取值的 Go 布局放在一起
// 光用 time.Parse 不够："15:04:05" 布局也接受 "9:04:05"
type rs时间格式 struct {
	pattern *regexp.Regexp
	s格式布局   string
}

func newRs时间格式(pattern string, s格式布局 string) *rs时间格式 {
	return &rs时间格式{
		pattern: regexp.MustCompile(pattern),
		s格式布局:   s格式布局,
	}
}

// newRs布局格式 derives the regex from a layout made of fixed-width numbers: each digit becomes \d, the rest stays as it is
// The layouts in this package are fixed-width; the tests check that each field's output passes its own check
//
// newRs布局格式 从位数固定的数字布局推出正则：每个数字换成 \d，其余字符原样保留
// 包里用的布局都是位数固定的；测试会核对每个字段格式化出来的文本都能过它自己的校验
func newRs布局格式(s格式布局 string) *rs时间格式 {
	var pattern strings.Builder
	pattern.WriteString("^")
	for _, c := range s格式布局 {
		if c >= '0' && c <= '9' {
			pattern.WriteString(`\d`)
		} else {
			pattern.WriteString(regexp.QuoteMeta(string(c)))
		}
	}
	pattern.WriteString("$")
	return newRs时间格式(pattern.String(), s格式布局)
}

func (f *rs时间格式) match(s string) bool {
	if !f.pattern.MatchString(s) {
		return false
	}
	if _, err := time.Parse(f.s格式布局, s); err != nil {
		return false
	}
	return true
}

// parse checks the shape, then parses the value in the zone
// parse 先查格式，再按时区解析取值
func (f *rs时间格式) parse(s string, z时区位置 *time.Location) (time.Time, error) {
	if !f.pattern.MatchString(s) {
		return time.Time{}, erero.Errorf("格式不符合 %s: %s", f.s格式布局, s)
	}
	v时间, err := time.ParseInLocation(f.s格式布局, s, z时区位置)
	if err != nil {
		return time.Time{}, erero.WithMessage(err, "取值不对")
	}
	return v时间, nil
}

// Match日期 validates if string is an existing date in format "2006-01-02"
//
// Match日期 验证字符串是不是 "2006-01-02" 格式的真实日期
func (rs *rs正则匹配时间字符串) Match日期(s标准日期 string) bool {
	return rs.v日期.match(s标准日期)
}

// Match时刻 validates if string is an existing time-of-moment in format "15:04:05"
//
// Match时刻 验证字符串是不是 "15:04:05" 格式的真实时刻
func (rs *rs正则匹配时间字符串) Match时刻(s标准时刻 string) bool {
	return rs.v时刻.match(s标准时刻)
}

// Match时间 validates if string is an existing datetime in format "2006-01-02 15:04:05"
//
// Match时间 验证字符串是不是 "2006-01-02 15:04:05" 格式的真实时间
func (rs *rs正则匹配时间字符串) Match时间(s标准时间 string) bool {
	return rs.v时间.match(s标准时间)
}

// Match日期分时 validates if string is an existing datetime without seconds in format "2006-01-02 15:04"
//
// Match日期分时 验证字符串是不是 "2006-01-02 15:04" 格式的真实时间（不带秒）
func (rs *rs正则匹配时间字符串) Match日期分时(s string) bool {
	return rs.v日期分时.match(s)
}

// Match日期零点 validates if string is an existing date at midnight in format "2006-01-02 00:00:00"
//
// Match日期零点 验证字符串是不是 "2006-01-02 00:00:00" 格式的真实日期零点
func (rs *rs正则匹配时间字符串) Match日期零点(s string) bool {
	return rs.v日期零点.match(s)
}

// Match八位数字日期 validates if string is an existing date in 8-digit format "20060102"
//
// Match八位数字日期 验证字符串是不是 "20060102" 格式的真实日期
func (rs *rs正则匹配时间字符串) Match八位数字日期(s string) bool {
	return rs.v八位数字日期.match(s)
}

// Match六位数字时刻 validates if string is an existing time-of-moment in 6-digit format "150405"
//
// Match六位数字时刻 验证字符串是不是 "150405" 格式的真实时刻
func (rs *rs正则匹配时间字符串) Match六位数字时刻(s string) bool {
	return rs.v六位数字时刻.match(s)
}

// Match八位数字时刻 validates if string is 8 digits: an existing time-of-moment "150405" and then 2 digits of the second's fraction
//
// Match八位数字时刻 验证字符串是不是 8 位数字：前 6 位是真实时刻 "150405"，后 2 位是秒以下的小数
func (rs *rs正则匹配时间字符串) Match八位数字时刻(s string) bool {
	return rs.v8位数字.MatchString(s) && rs.v六位数字时刻.match(s[:6])
}

// Match四位数字时分 validates if string is an existing time of minutes in 4-digit format "1504"
//
// Match四位数字时分 验证字符串是不是 "1504" 格式的真实时分
func (rs *rs正则匹配时间字符串) Match四位数字时分(s string) bool {
	return rs.v四位数字时分.match(s)
}

// Match年份 validates if string has 4 digits, such as "2006"
//
// Match年份 验证字符串是不是 4 位数年份 "2006"
func (rs *rs正则匹配时间字符串) Match年份(s四位年份 string) bool {
	return rs.v年份.match(s四位年份)
}
