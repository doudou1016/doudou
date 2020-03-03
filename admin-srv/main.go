package main

import (
	"doudou/admin-srv/handler"

	admin "doudou/admin-srv/proto/admin"

	"doudou/pkg/microplus"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/service"
)

func main() {

	/************************************/
	/********** New GRPC Service   ********/
	/************************************/
	// New Service
	service := microplus.NewService(
		service.Name("com.lcb123.srv.admin"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	admin.RegisterAdminHandler(service.Server(), new(handler.Admin))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("com.lcb123.srv.admin", service.Server(), new(subscriber.Admin))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
