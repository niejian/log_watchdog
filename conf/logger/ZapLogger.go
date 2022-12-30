package logger

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"org.code4fun/log/conf/global"
	"org.code4fun/log/utils"
)

// 初始化zap信息 https://www.cnblogs.com/guohewei/p/15256698.html
func InitLogger(conf ZapLogConfigs) error {
	logLevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}

	writeSync, err := getLogWriter(conf)
	if nil != err {
		return err
	}
	if logFmt := conf.LogFormat; len(logFmt) == 0 {
		conf.LogFormat = "logfmt"
	}
	encoder := getEncoder(conf)
	level, ok := logLevel[conf.LogLevel]
	if !ok {
		level = logLevel["info"]
	}

	core := zapcore.NewCore(encoder, writeSync, level)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	global.Log = logger.Sugar()
	return nil
}

func getEncoder(conf ZapLogConfigs) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// log 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 输出level序列化为全大写字符串，如 INFO DEBUG ERROR
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if conf.LogFormat == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)

}

func getLogWriter(conf ZapLogConfigs) (zapcore.WriteSyncer, error) {
	if exist := utils.FileIsExists(conf.LogPath); !exist {
		err := os.MkdirAll(conf.LogPath, os.ModePerm)
		if nil != err {
			return nil, err
		}
	}

	// 日志切割配置
	ZapLumberJack := &lumberjack.Logger{
		Filename:   filepath.Join(conf.LogPath, conf.LogFileName),
		MaxSize:    conf.LogFileMaxSize,
		MaxBackups: conf.LogFileMaxBackups,
		MaxAge:     conf.LogMaxAge,
		Compress:   conf.LogCompress,
	}
	if conf.LogStdout {
		return zapcore.
				NewMultiWriteSyncer(zapcore.AddSync(ZapLumberJack), zapcore.AddSync(os.Stdout)),
			nil
	} else {
		return zapcore.AddSync(ZapLumberJack), nil
	}

}
