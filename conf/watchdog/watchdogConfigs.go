package watchdog

type WatchdogConfigs struct {
	TraceInfos []TraceInfo `mapstructure:"traceInfo" json:"traceInfo" yaml:"trace-info"`
	Enable     bool        `mapstructure:"enable" json:"enable" yaml:"enable"`
}

type TraceInfo struct {
	LogPath   string   `mapstructure:"logPath" json:"logPath" yaml:"log-path"`
	ToUserIds string   `mapstructure:"toUserIds" json:"toUserIds" yaml:"to-user-ids"`
	Errs      []string `mapstructure:"errs" json:"errs" yaml:"errs"`
	Ignores   []string `mapstructure:"ignores" json:"ignores" yaml:"ignores"`
}
