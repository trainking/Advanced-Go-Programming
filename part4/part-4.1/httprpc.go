package main

import (
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"io"
)

// 名字
const HelloServiceName = "HelloService"

// 定义接口
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// 注册方法
func RegisterHelloService(h HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, h)
}

// 用结构体去实例化接口
type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}


func main() {
	RegisterHelloService(new(HelloService))
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter,r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		} {
			ReadCloser: r.Body,
			Writer: w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}