package config

import "os"

// TaskDownloadCompletedFallbackAddr 任务下载完成回调地址
var TaskDownloadCompletedFallbackAddr = os.Getenv("task_download_completed_fallback_addr")
