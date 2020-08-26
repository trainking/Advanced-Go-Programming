package main

// 用gopath引入注意关闭export GO111MODULE=off
import "./watch"
import (
	"net"
	"net/rpc"
	"log"
)

const WatchServiceName = "/path/to/WatchService"

func main() {
	var watchService watch.KVStoreServiceInterface
	watchService = watch.NewKVStoreService()
	rpc.RegisterName(WatchServiceName, watchService)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen tcp error : ", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error: ", err)
	}
	rpc.ServeConn(conn)
}