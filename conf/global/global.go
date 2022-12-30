package global

import "go.uber.org/zap"

// 定义全局变量
var (
	// use the SugaredLogger. It's 4-10x faster than other structured logging packages and includes both structured and printf-style APIs.
	Log *zap.SugaredLogger
)
