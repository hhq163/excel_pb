package main

import (
	"log"
	"os"
	"time"

	"github.com/hhq163/excel_pb"
	"github.com/hhq163/excel_pb/impl"
)

func main() {
	loginit()

	impl.ExcelToProto(excel_pb.GetExecpath()+"/input", excel_pb.GetExecpath()+"/output", 3)
	impl.ProtoToBytes(excel_pb.GetExecpath()+"/input", excel_pb.GetExecpath()+"/output")
	log.Println("create file success !")

}

//日志目录初始化
func loginit() {
	file := "./logs/" + "hub" + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
