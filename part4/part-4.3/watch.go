package main

import (
	"net/rpc"
	"fmt"
	"time"
	"log"
)

const WatchServiceName = "/path/to/WatchService"

func doClientWork(client *rpc.Client)  {
	go func ()  {
		var keyChanged string
		err := client.Call(WatchServiceName + ".Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	} ()

	err := client.Call(
		WatchServiceName + ".Set", [2]string{"abc", "abc-value"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 3)
}


func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	doClientWork(client)
}