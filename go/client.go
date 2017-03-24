package main

import (
	"fmt"

	proto "github.com/micro/examples/service/proto"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService(micro.Name("greeter.client"))

	greeter := proto.NewGreeterClient("greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Joaozito"})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Greeting)
}
