package mygin

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	"code.cloudfoundry.org/bytefmt"
)

var FuncMap = template.FuncMap{
	"toValMap": func(val interface{}) map[string]interface{} {
		return map[string]interface{}{
			"Value": val,
		}
	},
	"tf": func(t time.Time) string {
		return t.Format("01/02/2006 15:04:05")
	},
	"len": func(slice []interface{}) string {
		return strconv.Itoa(len(slice))
	},
	"safe": func(s string) template.HTML {
		return template.HTML(s) // #nosec
	},
	"tag": func(s string) template.HTML {
		return template.HTML(`<` + s + `>`) // #nosec
	},
	"stf": func(s uint64) string {
		return time.Unix(int64(s), 0).Format("01/02/2006 15:04")
	},
	"sf": func(duration uint64) string {
		return time.Duration(time.Duration(duration) * time.Second).String()
	},
	"sft": func(future time.Time) string {
		return time.Until(future).Round(time.Second).String()
	},
	"bf": func(b uint64) string {
		return bytefmt.ByteSize(b)
	},
	"ts": func(s string) string {
		return strings.TrimSpace(s)
	},
	"float32f": func(f float32) string {
		return fmt.Sprintf("%.3f", f)
	},
	"divU64": func(a, b uint64) float32 {
		if b == 0 {
			if a > 0 {
				return 100
			}
			return 0
		}
		if a == 0 {
			// 这是从未在线的情况
			return 0.00001 / float32(b) * 100
		}
		return float32(a) / float32(b) * 100
	},
	"div": func(a, b int) float32 {
		if b == 0 {
			if a > 0 {
				return 100
			}
			return 0
		}
		if a == 0 {
			// 这是从未在线的情况
			return 0.00001 / float32(b) * 100
		}
		return float32(a) / float32(b) * 100
	},
	"addU64": func(a, b uint64) uint64 {
		return a + b
	},
	"add": func(a, b int) int {
		return a + b
	},
	"TransLeftPercent": func(a, b float64) (n float64) {
		n, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", (100-(a/b)*100)), 64)
		if n < 0 {
			n = 0
		}
		return
	},
	"TransLeft": func(a, b uint64) string {
		if a < b {
			return "0B"
		}
		return bytefmt.ByteSize(a - b)
	},
	"TransClassName": func(a float64) string {
		if a == 0 {
			return "offline"
		}
		if a > 50 {
			return "fine"
		}
		if a > 20 {
			return "warning"
		}
		if a > 0 {
			return "error"
		}
		return "offline"
	},
	"UintToFloat": func(a uint64) (n float64) {
		n, _ = strconv.ParseFloat((strconv.FormatUint(a, 10)), 64)
		return
	},
	"dayBefore": func(i int) string {
		year, month, day := time.Now().Date()
		today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
		return today.AddDate(0, 0, i-29).Format("01/02")
	},
	"className": func(percent float32) string {
		if percent == 0 {
			return ""
		}
		if percent > 95 {
			return "good"
		}
		if percent > 80 {
			return "warning"
		}
		return "danger"
	},
}
