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
	return p.Client.Call(HelloServiceName + ".Hello", request, reply)
}

func main() {
	// client, err := rpc.Dial("tcp", "localhost:1234")
	client, err := DialHelloServiceClient("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	var reply string
	// err = client.Call(HelloServiceName + ".Hello", "hello, gg", &reply)
	err = client.Hello("hello,gg", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}