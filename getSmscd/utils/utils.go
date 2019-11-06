package utils

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/consul"
)

func GetGrpcService()micro.Service{
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.137.130:8500",
		}
	})
	grpcServer := grpc.NewService(
		micro.Registry(reg),
	)
	return grpcServer
}
