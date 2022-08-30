module test.go

go 1.17

// require github.com/hhq163/excel_pb v0.0.0-20220830071832-2809ae0dcc02
replace github.com/hhq163/excel_pb => ../../excel_pb

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/hhq163/excel_pb v0.0.0-00010101000000-000000000000 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	google.golang.org/protobuf v1.23.0 // indirect
)
