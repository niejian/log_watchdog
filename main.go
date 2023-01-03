package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
	"gopkg.in/yaml.v3"
	"org.code4fun/log/conf/global"
	"org.code4fun/log/conf/logger"
	"org.code4fun/log/core"
)

func init() {
	// 初始化，读取配置文件信息
	// 读取配置文件
	yamlData, err := ioutil.ReadFile("config.yaml")
	if nil != err {
		log.Fatal(err)
		panic("读取配置文件config.yaml失败")
	}

	err = yaml.Unmarshal(yamlData, &global.BaseConf)

	if nil != err {
		log.Fatal(err)
	}

	zapLogger, err := logger.InitLogger(global.BaseConf.LogFileConfig)
	if err != nil {
		log.Fatal(err)
	}
	// 初始化日志
	global.Log = zapLogger.Sugar()
	// 初始化本地缓存；过期时间30s，每10清除过期key
	global.LocalCache = cache.New(30*time.Second, 10*time.Second)

}

func main() {
	config := global.BaseConf.WatchdogConfig
	if !config.Enable {
		global.Log.Info("已关闭看门狗告警！")
		return
	}

	traceInfos := config.TraceInfos
	if len(traceInfos) == 0 {
		global.Log.Info("请配置traceInfos项")
		return
	}
	// 阻塞进程
	done := make(chan bool)
	for i := 0; i < len(traceInfos); i++ {
		traceInfo := traceInfos[i]
		core.TailLog(&traceInfo)

	}

	<-done
}
