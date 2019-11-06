package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	"microProject/getSmscd/handler"
	getSmscd "microProject/getSmscd/proto/getSmscd"
)

func main() {
	reg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"192.168.137.130:8500",
		}
	})
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.getSmscd"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Address("192.168.137.130:15002"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getSmscd.RegisterGetSmscdHandler(service.Server(), new(handler.GetSmscd))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
