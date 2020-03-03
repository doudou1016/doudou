package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	exam "exam-srv/proto/exam"
)

type Exam struct{}

func (e *Exam) Handle(ctx context.Context, msg *exam.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *exam.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
