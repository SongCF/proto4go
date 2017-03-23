proto4go
====================

proto generate tool for go, included protoc, protoc-gen-go.

1. generate `*.pb.go` by protoc and protoc-gen-go.
2. generate `msgcode.go` which define the map of code(int) to msg(string) and msg(string) to code(int), such as:
```
var Code = map[string]int{
	"heartbeat_req": 20061, // heart beat
	"heartbeat_ack": 20062, // heart beat response
}

var RCode = map[int]string{
	20061: "heartbeat_req", // heart beat
	20062: "heartbeat_ack", // heart beat response
}
```
3. generate `msgcode.csv` files


Already include protoc.exe(v3.2.0) and protoc-gen-go.exe(the last version 2017.3.22) in bin director.



# How to use

Not tested in Linux system

1. add `./bin` to path environment
2. `go install github.com/songcf/proto4go`
3. Run `proto4go.exe -i ./pb -o ./pb` to generate your proto files.



# sample

1. delete these files: `test/packet.pb.go`,`test/msgcode.go`,`test/msgcode.csv`
2. `go run main.go -i ./test -o ./test`
3. will generate these files: `packet.pb.go`,`msgcode.go`,`msgcode.csv`
