package global

import (
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"org.code4fun/log/conf"
)

// 定义全局变量
var (
	// use the SugaredLogger. It's 4-10x faster than other structured logging packages and includes both structured and printf-style APIs.
	Log        *zap.SugaredLogger
	BaseConf   *conf.BaseConfig
	LocalCache *cache.Cache
)
