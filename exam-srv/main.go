package main

import (
	"doudou/exam-srv/handler"

	exam "doudou/exam-srv/proto/exam"
	"doudou/pkg/microplus"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/service"
)

func main() {
	// New Service
	service := microplus.NewService(
		service.Name("com.lcb123.srv.exam"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	exam.RegisterExamHandler(service.Server(), new(handler.Exam))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("com.lcb123.srv.exam", service.Server(), new(subscriber.Exam))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
