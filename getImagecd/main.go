package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	"microProject/getImagecd/handler"
	getImagecd "microProject/getImagecd/proto/getImagecd"
)

func main() {
	// New Service
	reg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"192.168.137.130:8500",
		}
	})
	service := grpc.NewService(
		micro.Name("go.micro.srv.getImagecd"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Address("192.168.137.130:10000"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getImagecd.RegisterGetImagecdHandler(service.Server(), new(handler.GetImagecd))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
