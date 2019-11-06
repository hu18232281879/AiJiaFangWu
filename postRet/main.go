package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	"microProject/postRet/handler"
	postRet "microProject/postRet/proto/postRet"
)

func main() {
	// New Service
	config:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"192.168.137.130:8500",
		}
	})
	service := grpc.NewService(
		micro.Name("go.micro.srv.postRet"),
		micro.Version("latest"),
		micro.Registry(config),
		micro.Address("192.168.137.130:10004"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	postRet.RegisterPostRetHandler(service.Server(), new(handler.PostRet))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
