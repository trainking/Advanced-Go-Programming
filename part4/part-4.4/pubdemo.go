package main

import (
	pubsub "github.com/docker/docker/pkg/pubsub"
	"time"
	"strings"
	"fmt"
)

func main() {
	p := pubsub.NewPublisher(100 * time.Millisecond, 10)
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})
	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})
	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("docker: https://www.docker.org")
	time.Sleep(1 * time.Second)
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go func() {
		fmt.Println("golang topic:", <-golang)
		ch1 <- true
	}()
	go func() {
		fmt.Println("docker topic:", <-docker)
		ch2 <- true
	}()
	<-ch1
	<-ch2
}