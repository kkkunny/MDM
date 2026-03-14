package db

import (
	"fmt"
	"regexp"

	"github.com/duke-git/lancet/v2/convertor"
	stlos "github.com/kkkunny/stl/os"
	stlval "github.com/kkkunny/stl/value"

	"github.com/kkkunny/MDM/config"
)

var stlLogger = config.Logger.NewGroup("SQL_LITE")

type customLogger struct{}

func (l *customLogger) Printf(format string, v ...interface{}) {
	var logFn func(a ...any) error
	var content string
	var pos []stlos.Frame
	switch format {
	case "%s\n[info] ":
		logFn = stlLogger.Info
		content = convertor.ToString(v[0])
	case "%s\n[warn] ":
		logFn = stlLogger.Warn
		content = convertor.ToString(v[0])
	case "%s\n[error] ":
		logFn = stlLogger.Error
		content = convertor.ToString(v[0])
	case "%s\n[%.3fms] [rows:%v] %s":
		logFn = stlLogger.Trace
		content = convertor.ToString(v[3])
		matches := regexp.MustCompile(`(.+)?:(\d+)`).FindStringSubmatch(convertor.ToString(v[0]))
		if len(matches) > 0 {
			pos = append(pos, stlos.NewFrame(convertor.ToString(matches[1]), int(stlval.IgnoreWith(convertor.ToInt(matches[2])))))
		}
	case "%s %s\n[%.3fms] [rows:%v] %s":
		logFn = stlLogger.Trace
		content = convertor.ToString(v[4])
	default:
		logFn = stlLogger.Info
		content = fmt.Sprintf(format, v...)
	}
	stlval.Ignore(logFn(content, stlLogger.WithDisplayPosition(pos...)))
}
