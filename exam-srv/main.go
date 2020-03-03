package main

import (
	"exam-srv/handler"
	"exam-srv/subscriber"
	"time"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	exam "exam-srv/proto/exam"
)

func main() {
	/************************************/
	/********** 服务发现  etcd   ********/
	/************************************/
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:2379",
		}
		op.Timeout = 5 * time.Second
	})
	/************************************/
	/********** New GRPC Service   ********/
	/************************************/
	// New Service
	service := micro.NewService(
		micro.Name("com.lcb123.srv.exam"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*15),      //重新注册时间
		micro.RegisterInterval(time.Second*10), //注册过期时间
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	exam.RegisterExamHandler(service.Server(), new(handler.Exam))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.lcb123.srv.exam", service.Server(), new(subscriber.Exam))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
