package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
	"org.code4fun/log/conf/global"
	"org.code4fun/log/conf/logger"
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
	global.Log = zapLogger.Sugar()

}

func main() {
	global.Log.Infof("初始化成功。。。。")
}
