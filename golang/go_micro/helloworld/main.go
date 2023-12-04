package main

import (
	"context"
	"fmt"
	"log"

	pb "go_micro/helloworld/proto/proto"

	"go-micro.dev/v4"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Msg = "Hello " + req.Name
	fmt.Printf("[] recv %s, rsp %s\n", req.Name, rsp.Msg)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)
	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
