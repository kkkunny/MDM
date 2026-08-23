package config

import (
	"os"

	stllog "github.com/kkkunny/stl/log"
	stlval "github.com/kkkunny/stl/value"
)

var Logger *stllog.Logger

func init() {
	var level stllog.Level
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		level = stllog.LevelDebug
	case "TRACE":
		level = stllog.LevelTrace
	case "INFO":
		level = stllog.LevelInfo
	case "WARN":
		level = stllog.LevelWarn
	case "KEYWORD":
		level = stllog.LevelKeyword
	case "ERROR":
		level = stllog.LevelError
	case "PANIC":
		level = stllog.LevelPanic
	default:
		level = stlval.If(Release, stllog.LevelInfo, stllog.LevelDebug)
	}
	Logger = stllog.New(os.Stdout, level)
	Logger.SetDefaultConfig(Logger.WithDisplayLevel().WithDisplayTime().WithDisplayPosition().WithDisplayColor().WithDisplayGroup())
}
