package main

import (
	"admin-srv/handler"
	"admin-srv/subscriber"
	"time"

	admin "admin-srv/proto/admin"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/service/grpc"
)

func main() {
	/************************************/
	/********** 服务发现  etcd   ********/
	/************************************/
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
		op.Timeout = 5 * time.Second
	})
	/************************************/
	/********** New GRPC Service   ********/
	/************************************/
	// New Service
	service := grpc.NewService(
		micro.Name("com.lcb123.srv.admin"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*15),      //重新注册时间
		micro.RegisterInterval(time.Second*10), //注册过期时间
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	admin.RegisterAdminHandler(service.Server(), new(handler.Admin))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.lcb123.srv.admin", service.Server(), new(subscriber.Admin))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
