package impl

//将excel生成proto，再生成pb文件
func ExcelToPb(input, output string, protoVer int32) {
	ExcelToProto(input, output, protoVer)
	ProtoToBytes(input, output)
}
