package main


import (
	"net"
	"net/rpc/jsonrpc"
	"net/rpc"
	"log"
)

// 名字
const HelloServiceName = "/path/to/HelloService"

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
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listene tcp error : ", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
}