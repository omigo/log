package log

// 可以用这些串和日期、时间（包含毫秒数）任意组合，拼成各种格式的日志，如 csv/json/xml
const (
	LevelToken   string = "info"
	TagToken            = "tag"
	PathToken           = "/go/src/github.com/gotips/log/examples/main.go"
	PackageToken        = "github.com/gotips/log/examples/main.go"
	ProjectToken        = "examples/main.go"
	FileToken           = "main.go"
	LineToken    int    = 88
	MessageToken string = "message"
)

// DefaultFormat 默认日志格式
const DefaultFormat = "2006-01-02 15:04:05 info examples/main.go:88 message"

// DefaultFormatTag 默认日志格式带标签
const DefaultFormatTag = "2006-01-02 15:04:05 tag info examples/main.go:88 message"
