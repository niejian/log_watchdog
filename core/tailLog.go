package core

import (
	"strings"
	"time"

	"github.com/hpcloud/tail"
	"org.code4fun/log/conf/global"
	"org.code4fun/log/conf/watchdog"
	"org.code4fun/log/utils"
)

var (
	ErrorTag = "ERROR"
)

/**
对文件执行tail操作
*/
func TailLog(traceInfo *watchdog.TraceInfo) {
	tailFile, err := initTail(traceInfo.LogPath)
	if nil != err {
		global.Log.Errorf("文件：%s 初始化tail失败")
		return
	}
	var errMsg string
	// 执行类似tail -f 读取行
	for line := range tailFile.Lines {
		newLine := line.Text
		if len(newLine) == 0 {
			continue
		}

		// 是否是错误日志行
		// 错误日志行有以下特点：1. 含有ERROR；2. 以tab开头
		isErrLine := false

		// 日期开头含有ERROR
		if utils.IsDatePrefix(newLine) && strings.Contains(newLine, ErrorTag) {
			isErrLine = true
		}
		// tab开头
		if !utils.IsDatePrefix(newLine) && strings.HasPrefix(newLine, "\t") {
			isErrLine = true
		}

		// 判断是否含有关键字
		errs := traceInfo.Errs
		ingnores := traceInfo.Ignores
		if len(errs) == 0 {
			global.Log.Errorf("应用：%s, 请填写告警异常", traceInfo.AppName)
			continue
		}
		if isErrLine &&
			utils.StrInArr(errs, newLine) &&
			!utils.StrInArr(ingnores, newLine) {
			errMsg += newLine
		}

		time.AfterFunc(500*time.Millisecond, func() {

		})

	}

}

func initTail(fileName string) (*tail.Tail, error) {
	return tail.TailFile(fileName, tail.Config{
		ReOpen:    true,
		Follow:    true,                                 // 实时跟踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取。 whence，从哪开始：0从头，1当前，2末尾
		MustExist: false,                                // 如果文件不存在，是否推出程序，false是不退出
		Poll:      true,
	})
}
