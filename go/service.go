package main

import (
	"fmt"

	proto "github.com/micro/examples/service/proto"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	// Init will parse the command line flags. Any flag set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	service.Init()

	// Register Handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the Server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
