package impl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

//生成proto go版本目标文件
func GenProto(outPath string) {
	files, err := ioutil.ReadDir(outPath)
	if err != nil {
		log.Fatalln("outPath file error: ", outPath, err.Error())
	}
	for _, file := range files {
		fileName := file.Name()
		paramArr := make([]string, 0)
		paramArr = append(paramArr, fmt.Sprintf("--proto_path=%s", outPath))
		paramArr = append(paramArr, fmt.Sprintf("--gofast_out=%s", outPath))
		paramArr = append(paramArr, fmt.Sprintf("%s/%s", outPath, fileName))

		log.Println("protoStr " + strings.Join(paramArr, " "))
		cmd := exec.Command("protoc", paramArr...)

		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalln("proto 生成工具出错, err=", err.Error(), ",stderr=", stderr.String())
			fmt.Println("proto 生成工具出错 ", err.Error())
			return
		}
	}

}
