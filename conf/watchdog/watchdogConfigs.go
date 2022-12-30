package watchdog

type WatchdogConfigs struct {
	TraceInfo TraceInfo `mapstructure:"traceInfo" json:"traceInfo" yaml:"traceInfo"`
	Enable    bool      `mapstructure:"enable" json:"enable" yaml:"enable"`
}

type TraceInfo struct {
	LogPath   string   `mapstructure:"logPath" json:"logPath" yaml:"logPath"`
	ToUserIds string   `mapstructure:"toUserIds" json:"toUserIds" yaml:"toUserIds"`
	Errs      []string `mapstructure:"errs" json:"errs" yaml:"errs"`
	Ignores   []string `mapstructure:"ignores" json:"ignores" yaml:"ignores"`
}
