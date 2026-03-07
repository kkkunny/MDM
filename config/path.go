package config

import (
	"path/filepath"

	stlval "github.com/kkkunny/stl/value"
)

const (
	XLDownloadDirByXL = "/迅雷下载"      // 迅雷视角的迅雷下载目录
	QBDownloadDirByQB = "/downloads" // qb视角的qb下载目录
)

var (
	DownloadDir         = stlval.Ternary(Release, "/downloads", "/mnt/data/downloads") // 迅雷下载目录
	CompleteDownloadDir = filepath.Join(DownloadDir, "complete")
	XLBtDir             = filepath.Join(DownloadDir, ".bt") // 种子文件存放目录
)
