package log

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

type record struct {
	Date, Time string
	Level      string
	File       string
	Line       int
	Message    string
}

// DefaultFormat 默认日志格式
const DefaultFormat = "2006-01-02 15:04:05 info examples/main.go:88 message"

// format tokens, 和 日期、时间组成日志格式
const (
	LevelToken   string = "info"
	PathToken           = "/go/src/github.com/gotips/log/examples/main.go"
	PackageToken        = "github.com/gotips/log/examples/main.go"
	ProjectToken        = "examples/main.go"
	FileToken           = "main.go"
	LineToken    int    = 88
	MessageToken string = "message"
)

// Standard 日志输出基本实现
type Standard struct {
	mu  sync.Mutex // ensures atomic writes; protects the following fields
	out io.Writer  // destination for output

	format    string // log format
	pattern   string // log template
	tpl       *template.Template
	prefixLen int
	dateFmt   string
	timeFmt   string
}

// NewStandard creates a new Logger.   The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line.
// The flag argument defines the logging properties.
func NewStandard(out io.Writer, format string) *Standard {
	std := &Standard{out: out}
	std.SetFormat(format)

	return std
}

// ChangeWriter sets the output destination for the logger.
func (s *Standard) ChangeWriter(w io.Writer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.out = w
}

// SetFormat set output format for the printer
func (s *Standard) SetFormat(format string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// println(format)
	s.format = format

	s.pattern = format

	s.prefixLen = calculatePrefixLen(format, 3)

	// 提取出日期和时间的格式化模式字符串
	s.dateFmt, s.timeFmt = extactDateTimeFormat(format)

	s.pattern = strings.Replace(s.pattern, LevelToken, "{{ .Level }}", -1)
	s.pattern = strings.Replace(s.pattern, PathToken, "{{ .File }}", -1)
	s.pattern = strings.Replace(s.pattern, PackageToken, "{{ .File }}", -1)
	s.pattern = strings.Replace(s.pattern, ProjectToken, "{{ .File }}", -1)
	s.pattern = strings.Replace(s.pattern, FileToken, "{{ .File }}", -1)
	s.pattern = strings.Replace(s.pattern, strconv.Itoa(LineToken), "{{ .Line }}", -1)
	s.pattern = strings.Replace(s.pattern, MessageToken, "{{ .Message }}", -1)

	// println(s.dateFmt, s.timeFmt)

	if s.dateFmt != "" {
		s.pattern = strings.Replace(s.pattern, s.dateFmt, "{{ .Date }}", -1)
	}
	if s.timeFmt != "" {
		s.pattern = strings.Replace(s.pattern, s.timeFmt, "{{ .Time }}", -1)
	}
	s.pattern += "\n"

	s.tpl = template.Must(template.New("record").Parse(s.pattern))
}

// Print 打印日志
func (s *Standard) Print(l Level, m string) error {
	r := record{
		Level:   l.String(),
		Message: strings.TrimRight(m, "\n"),
	}

	if s.dateFmt != "" {
		now := time.Now() // get this early.
		r.Date = now.Format(s.dateFmt)
		if s.timeFmt != "" {
			r.Time = now.Format(s.timeFmt)
		}
	}

	if s.prefixLen > -1 {
		var ok bool
		_, r.File, r.Line, ok = runtime.Caller(3) // expensive
		if ok {
			r.File = r.File[s.prefixLen:]
		} else {
			r.File = "???"
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.tpl.Execute(s.out, r)
	return err
}

func extactDateTimeFormat(format string) (dateFmt, timeFmt string) {
	// 算法：
	// 找出两个字符串不同的部分，
	// 如果有两处不同，一个是日期模式，一个是时间模式，
	// 如果只有一个，那么只有日期或者只有时间，无关紧要，
	// 如果都相同，那么日志里没有时间，
	// 如果有三处以上不同，说明格式配置错误
	// TODO 有 bug，比如日期格式如果配置成 2006-1-2，那么两个串长度不等

	t, _ := time.ParseInLocation("2006-1-2 3:4:5.000000000", "1991-2-1 1:1:1.111111111", time.Local)
	contrast := t.Format(format)

	// println(format)
	// println(contrast)

	idxs := [10]int{}
	start := -1
	for i, l, same := 0, len(format), true; i < l; i++ {
		if start > 4 {
			panic(fmt.Sprintf("format string error at `%s`", format[i-1:]))
		}

		// fmt.Printf("%c %c %d %d\n", format[i], contrast[i], idxs, start)

		if format[i] != contrast[i] {
			if same {
				start++
				// 如果之前都是相同的，这个开始不同，那么这个就是起始位置
				idxs[start] = i
				same = false
				// println(i, diff, start, idxs[start])
				start++
			}

			idxs[start] = i + 1 // 下一个有可能是结束位置

			continue
		}

		// 如果是 空格、-、:、. ，那么它不一定是结束位置
		if format[i] == '-' || format[i] == ' ' || format[i] == ':' || format[i] == '.' {
			// 如果这些字符后面是 0（如 2006-01-02），跳过 0
			if i+1 < l && format[i+1] == '0' && contrast[i+1] == '0' {
				i++
			}
			continue
		}

		same = true
	}

	if start != -1 && start != 1 && start != 3 {
		// 正常情况是不可能到这里的，如果到这里，说明算法写错了
		panic(fmt.Sprintf("parse error %d", start))

	} else {
		if start > 0 {
			dateFmt = format[idxs[0]:idxs[1]]
			if start == 3 {
				timeFmt = format[idxs[2]:idxs[3]]
			}
		}
	}

	return dateFmt, timeFmt
}

func calculatePrefixLen(format string, skip int) int {
	// 格式中不包含文件路径
	if !strings.Contains(format, "main.go") {
		return -1
	}

	_, file, _, _ := runtime.Caller(skip)

	// file with absolute path
	if strings.Contains(format, PathToken) {
		return 0
	}

	// file with package name
	if strings.Contains(format, PackageToken) {
		return strings.Index(file, "/src/") + 5
	}

	// file with project path
	if strings.Contains(format, ProjectToken) {
		// remove /<GOPATH>/src/
		prefixLen := strings.Index(file, "/src/") + 5
		file = file[prefixLen:]

		// remove github.com/
		if strings.HasPrefix(file, "github.com/") {
			prefixLen += 11
			file = file[11:]

			// remove github user or org name
			if i := strings.Index(file, "/"); i >= 0 {
				prefixLen += i + 1
				file = file[i+1:]

				// remove project name
				if i := strings.Index(file, "/"); i >= 0 {
					prefixLen += i + 1
				}
			}
		}

		return prefixLen
	}

	// file only
	if strings.Contains(format, FileToken) {
		return strings.LastIndex(file, "/") + 1
	}

	return -1
}
