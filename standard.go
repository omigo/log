package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

type record struct {
	Date, Time string
	Tag        string
	Level      string
	File       string
	Line       int
	Message    string
	Stack      []byte
}

// Standard 日志输出基本实现
type Standard struct {
	mu  sync.Mutex // ensures atomic writes; protects the following fields
	out io.Writer  // destination for output

	tpl       *template.Template
	prefixLen int
	dateFmt   string
	timeFmt   string

	defaultOne bool
}

// NewStandard 返回标准实现
func NewStandard(out io.Writer, format string) *Standard {
	std := &Standard{out: out}

	// hack 如果用户不调用 SetFormat，直接用，那么也能找到主函数（main，实际是 init 函数）的所在的文件
	std.prefixLen = -5

	std.SetFormat(format)
	return std
}

// SetWriter 改变输出流
func (s *Standard) SetWriter(w io.Writer) {
	s.mu.Lock()
	s.out = w
	s.mu.Unlock()
}

// SetFormat 改变日志格式
func (s *Standard) SetFormat(format string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	skip := 3
	if s.prefixLen == -5 {
		skip = 5
	}
	s.prefixLen = CalculatePrefixLen(format, skip)

	s.dateFmt, s.timeFmt = ExtactDateTime(format)

	pattern := parseFormat(format, s.prefixLen, s.dateFmt, s.timeFmt)

	s.tpl = template.Must(template.New("record").Parse(pattern))
}

// Tprintf 打印日志
func (s *Standard) Tprintf(v, l Level, tag string, format string, m ...interface{}) {
	if v > l {
		return
	}

	if tag == "" {
		tag = "-"
	}
	r := record{
		Level: l.String(),
		Tag:   tag,
	}

	if s.dateFmt != "" {
		now := time.Now() // get this early.
		r.Date = now.Format(s.dateFmt)
		if s.timeFmt != "" {
			r.Time = now.Format(s.timeFmt)
		}
	}

	if format == "" {
		r.Message = fmt.Sprint(m...)
	} else {
		r.Message = fmt.Sprintf(format, m...)
	}
	r.Message = strings.TrimRight(r.Message, "\n")

	if s.prefixLen > -1 {
		var ok bool
		_, r.File, r.Line, ok = runtime.Caller(2) // expensive
		if ok && s.prefixLen < len(r.File) {
			r.File = r.File[s.prefixLen:]
		} else {
			r.File = "???"
		}
	}

	var buf []byte
	if l == StackLevel {
		buf = make([]byte, 1024*1024)
		n := runtime.Stack(buf, true)
		buf = buf[:n]
	}

	s.mu.Lock()
	defer func() {
		s.mu.Unlock()

		if l == PanicLevel {
			panic(m)
		}

		if l == FatalLevel {
			os.Exit(-1)
		}
	}()

	s.tpl.Execute(s.out, r)
	s.out.Write([]byte("\n"))

	if l == StackLevel {
		s.out.Write(buf)
	}
}

// 格式解析，把格式串替换成 token 串
func parseFormat(format string, prefixLen int, dateFmt, timeFmt string) (pattern string) {
	// 顺序最好不要变，从最长的开始匹配
	pattern = strings.Replace(format, PathToken, "{{ .File }}", -1)
	pattern = strings.Replace(pattern, PackageToken, "{{ .File }}", -1)
	pattern = strings.Replace(pattern, ProjectToken, "{{ .File }}", -1)
	pattern = strings.Replace(pattern, FileToken, "{{ .File }}", -1)
	pattern = strings.Replace(pattern, TagToken, "{{ .Tag }}", -1)
	pattern = strings.Replace(pattern, LevelToken, "{{ .Level }}", -1)
	pattern = strings.Replace(pattern, strconv.Itoa(LineToken), "{{ .Line }}", -1)
	pattern = strings.Replace(pattern, MessageToken, "{{ .Message }}", -1)

	// 提取出日期和时间的格式化模式字符串
	if dateFmt != "" {
		pattern = strings.Replace(pattern, dateFmt, "{{ .Date }}", -1)
	}
	if timeFmt != "" {
		pattern = strings.Replace(pattern, timeFmt, "{{ .Time }}", -1)
	}
	return pattern
}
