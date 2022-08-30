package main

import (
	"log"
	"os"
	"time"

	"github.com/hhq163/excel_pb/impl"
)

func main() {
	loginit()

	impl.ExcelToProto()
	impl.ProtoToBytes()
	log.Println("create file success !")

}

func loginit() {
	file := "./logs/" + "hub" + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
