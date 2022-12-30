package logger

type ZapLogConfigs struct {
	LogLevel          string `mapstructure:"logLevel" json:"logLevel" yaml:"logLevel"`    // 日志级别
	LogFormat         string `mapstructure:"logFormat" json:"logFormat" yaml:"logFormat"` // 日志格式 json, logfmt
	LogPath           string `mapstructure:"logPath" json:"logPath" yaml:"logPath"`       // 日志保存目录
	LogFileName       string `mapstructure:"logFileName" json:"logFileName" yaml:"logFileName"`
	LogFileMaxSize    int    `mapstructure:"logFileMaxSize" json:"logFileMaxSize" yaml:"logFileMaxSize"`          // 【日志分割】单个日志文件最多存储量 单位(mb)
	LogFileMaxBackups int    `mapstructure:"logFileMaxBackups" json:"logFileMaxBackups" yaml:"logFileMaxBackups"` // 【日志分割】日志备份文件最多数量
	LogMaxAge         int    `mapstructure:"logMaxAge" json:"logMaxAge" yaml:"logMaxAge"`                         // 日志保留时间，单位: 天 (day)
	LogCompress       bool   `mapstructure:"logCompress" json:"logCompress" yaml:"logCompress"`                   // 是否压缩日志
	LogStdout         bool   `mapstructure:"logStdout" json:"logStdout" yaml:"logStdout"`                         // 是否输出到控制台
}
