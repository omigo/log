package log

import "testing"

func TestCalculatePrefixLen(t *testing.T) {
	{
		format := `{"level": "info", "line": 88, "log": "message"}`
		prefixLen := calculatePrefixLen(format, 1)
		if prefixLen != -1 {
			t.Error("prefix len error")
			t.FailNow()
		}
	}
	{
		format := `{"level": "info", "file": "/go/src/github.com/gotips/log/examples/main.go", "line":88, "log": "message"}`
		prefixLen := calculatePrefixLen(format, 1)
		if prefixLen != 0 {
			t.Error("prefix len error")
			t.FailNow()
		}
	}
	{
		format := `{"level": "info", "file": "github.com/gotips/log/examples/main.go", "line":88, "log": "message"}`
		prefixLen := calculatePrefixLen(format, 1)
		if prefixLen != len("/opt/gowork/src/") {
			t.Error("prefix len error")
			t.FailNow()
		}
	}

	{
		format := `{"level": "info", "file": "examples/main.go", "line":88, "log": "message"}`
		prefixLen := calculatePrefixLen(format, 1)
		if prefixLen != len("/opt/gowork/src/github.com/gotips/log/") {
			t.Error("prefix len error")
			t.FailNow()
		}
	}
	{
		format := `{"level": "info", "file": "main.go", "line":88, "log": "message"}`
		prefixLen := calculatePrefixLen(format, 1)
		if prefixLen != len("/opt/gowork/src/github.com/gotips/log/") {
			t.Error("prefix len error")
			t.FailNow()
		}
	}
}

func TestExtactDateTimeFormat(t *testing.T) {
	{
		format := `{"level": "info", "file": "log/main.go", "line":88, "log": "message"}`
		dateFmt, timeFmt := extactDateTimeFormat(format)
		if dateFmt != "" && timeFmt != "" {
			t.Error("format parse error")
			t.FailNow()
		}
	}

	{
		format := `{"datetime": "2006-01-02 15:04:05.999999999", "level": "info", "file": "log/main.go", "line":88, "log": "message"}`
		dateFmt, timeFmt := extactDateTimeFormat(format)
		if dateFmt != "2006-01-02 15:04:05.999999999" && timeFmt != "" {
			t.Error("data time format parse error")
			t.FailNow()
		}
	}

	{
		format := `{"date": "2006-01-02", "time": "15:04:05.999999999", "level": "info", "file": "log/main.go", "line":88, "log": "message"}`
		dateFmt, timeFmt := extactDateTimeFormat(format)
		if dateFmt != "2006-01-02" && timeFmt != "15:04:05.999999999" {
			t.Error("data time format parse error")
			t.FailNow()
		}
	}

	// 测试 日期模式不能重复出现在 format 中，不能判定是模式还是固定字符串
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("must panic, but not")
				t.FailNow()
			}
		}()

		// 有两个 2006 ，会出错
		format := `{"date": "2006-01-02", "time": "15:04:05.999999999", "traceID": "2006" "level": "info", "file": "log/main.go", "line":88, "log": "message"}`
		dateFmt, timeFmt := extactDateTimeFormat(format)
		if dateFmt != "2006-01-02" && timeFmt != "15:04:05.999999999" {
			t.Error("data time format parse error")
		}
	}()
}
