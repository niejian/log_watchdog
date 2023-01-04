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

// 判断字符串是否是日期时间戳开头
func IsDatePrefix(line string) bool {
	r := []rune(line)
	newLine20Prefix := string(r[0:19])
	pattern := "\\d{4}\\-\\d{2}\\-\\d{2}\\s\\d{2}:\\d{2}:\\d{2}"
	match, _ := regexp.Match(pattern, []byte(newLine20Prefix))
	return match
}

// 判断数组中的元素是否在目标字符串中
func StrInArr(datas []string, target string) bool {
	if len(datas) == 0 && len(target) == 0 {
		return true
	}
	for _, data := range datas {
		if strings.Contains(target, data) {
			return true
		}
	}
	return false
}
