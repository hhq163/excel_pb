package impl

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

//生成proto文件
func GenProto(outPath string) {
	protoStr := fmt.Sprintf(" --proto_path=%s --gofast_out=%s %s/*.proto", outPath, outPath, outPath)
	log.Println(protoStr)
	cmd := exec.Command("protoc", protoStr)

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
