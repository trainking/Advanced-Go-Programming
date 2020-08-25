package main

import (
	"net/rpc"
	"log"
	"fmt"
)

// 名字
const HelloServiceName = "/path/to/HelloService"

// 定义接口
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

type HelloServiceClient struct {
	Client *rpc.Client
}

func DialHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error{
	// 异步调用实现
	helloCall := p.Client.Go(HelloServiceName + ".Hello", request, reply, nil)
	// 这里干点别的, 因为没有等到数据，所以 reply 里面是个空
	fmt.Println("太阳出来唱山歌......->", *reply)

	// 通道获取到数据
	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func main() {
	client, err := DialHelloServiceClient("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	var reply string
	err = client.Hello("hello,gg", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("finaly", reply)
}