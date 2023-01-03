package core

import (
	"github.com/hpcloud/tail"
	"org.code4fun/log/conf/global"
	"org.code4fun/log/conf/watchdog"
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
