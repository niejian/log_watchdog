package logger

type ZapLogConfigs struct {
	LogLevel          string // 日志级别
	LogFormat         string // 日志格式 json, logfmt
	LogPath           string // 日志保存目录
	LogFileName       string
	LogFileMaxSize    int  // 【日志分割】单个日志文件最多存储量 单位(mb)
	LogFileMaxBackups int  // 【日志分割】日志备份文件最多数量
	LogMaxAge         int  // 日志保留时间，单位: 天 (day)
	LogCompress       bool // 是否压缩日志
	LogStdout         bool // 是否输出到控制台
}
