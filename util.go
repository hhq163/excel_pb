package excel_pb

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// 获取当前程序运行目录
func GetExecpath() string {
	execpath, _ := os.Executable() // 获得程序路径
	path := filepath.Dir(execpath)
	return strings.Replace(path, "\\", "/", -1)
}

//获取文件名
func getFileName(s string) string {
	arr := strings.Split(s, "_")
	fileName := arr[0]

	return fileName
}

//追加文件内容到末尾
func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}

	return err
}

//ToInt64 string to int64
func ToInt64(str string) int64 {
	if str == "" {
		return 0
	}
	id, _ := strconv.ParseInt(str, 10, 64)
	return id
}

//ToFloat string to float64
func ToFloat(str string) float64 {
	if str == "" {
		return 0
	}
	id, _ := strconv.ParseFloat(str, 64)
	return id
}

//ToInt32 string to int32
func ToInt32(str string) int32 {
	if str == "" {
		return 0
	}
	id, _ := strconv.Atoi(str)
	return int32(id)
}
