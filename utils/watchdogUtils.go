package utils

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func FileIsExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 字符串中是否含有数字
func HasDigital(str string) bool {
	pattern := "\\d{4}"
	match, _ := regexp.Match(pattern, []byte(str))
	return match
}
