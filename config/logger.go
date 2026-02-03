package config

import (
	stllog "github.com/kkkunny/stl/log"
)

var (
	Logger     = stllog.Default(!Release)
	XLLogger   = Logger.NewGroup("XUNLEI")
	HttpLogger = Logger.NewGroup("HTTP")
)
