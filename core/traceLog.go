package core

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"org.code4fun/log/conf/global"
	"org.code4fun/log/conf/watchdog"
	"org.code4fun/log/utils"
)

var fileMap sync.Map

// 日志追踪
func TraceLogFile(traceInfo *watchdog.TraceInfo) {
	done := make(chan bool)
	logFilePath := traceInfo.LogPath
	appName := traceInfo.AppName
	// 判断日志路径是否存在
	if exist := utils.FileIsExists(logFilePath); !exist {
		global.Log.Errorf("应用： %s，日志文件%s 不存在，请修改配置 %s", appName, logFilePath)

		return
	}
	global.Log.Infof("开始监听应用： %s，日志：%s", appName, logFilePath)
	watcher, err := fsnotify.NewWatcher()
	if nil != err {
		global.Log.Errorf("文件%s 监听失败，err： %v", logFilePath, err)
		return
	}

	defer func() {
		if err := recover(); nil != err {
			watcher.Close()
			global.Log.Errorf("文件追踪异常，%v", err)
		}
	}()
	watcher.Add(logFilePath)
	go func() {

		for {
			select {
			// 监听文件变化
			case event := <-watcher.Events:
				global.Log.Infof("fs event %s", event.Op)
				// 变化的文件名
				opName := event.Name
				if event.Op&fsnotify.Write == fsnotify.Write {
					// 判断文件是不是日志文件, .log或.out 结尾
					if !strings.HasSuffix(opName, ".log") &&
						!strings.HasSuffix(opName, ".out") {
						global.Log.Infof("文件：%s, 不是日志文件", opName)
						continue
					}
					// 判断文件是否是归档文件，文件名一般是含有日期，数字的
					traceFileNameArr := strings.Split(opName, string(filepath.Separator))
					// 获取追踪日志名称
					traceFileName := traceFileNameArr[len(traceFileNameArr)-1]
					if utils.HasDigital(traceFileName) {
						global.Log.Infof("文件：%s, 可能是归档文件，不处理", opName)
						continue
					}
					_, ok := fileMap.Load(logFilePath)
					if ok {
						global.Log.Infof("文件：%s 已被监听", logFilePath)
						continue
					}
					defer func() {
						err := recover()
						if nil != err {
							global.Log.Errorf("tail file: %s 失败", logFilePath)
						}
					}()
					go TailLog(traceInfo)

				}
			}
		}
	}()

	<-done
}
