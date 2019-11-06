package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	"microProject/getUserInfo/handler"
	getUserInfo "microProject/getUserInfo/proto/getUserInfo"
)

func main() {
	config:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"192.168.137.130:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getUserInfo"),
		micro.Version("latest"),
		micro.Registry(config),
		micro.Address("192.168.137.130:15100"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getUserInfo.RegisterGetUserInfoHandler(service.Server(), new(handler.GetUserInfo))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
