package global

import (
	"org.code4fun/log/conf/logger"
	"org.code4fun/log/conf/watchdog"
)

type BaseConfig struct {
	WatchdogConfig watchdog.WatchdogConfigs `mapstructure:"watchdogConfig" json:"watchdogConfig" yaml:"watchdogConfig"`
	LogFileConfig  logger.ZapLogConfigs     `mapstructure:"logFileConfig" json:"logFileConfig" yaml:"logFileConfig"`
}
