package conf

import (
	"org.code4fun/log/conf/logger"
	"org.code4fun/log/conf/watchdog"
)

type BaseConfig struct {
	WatchdogConfig watchdog.WatchdogConfigs `mapstructure:"watchdogConfig" json:"watchdogConfig" yaml:"watchdog-config"`
	LogFileConfig  logger.ZapLogConfigs     `mapstructure:"logFileConfig" json:"logFileConfig" yaml:"log-file-config"`
}
