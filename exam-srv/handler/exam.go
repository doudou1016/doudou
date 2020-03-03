package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	pb "doudou/admin-srv/proto/admin"
	exam "doudou/exam-srv/proto/exam"
	"doudou/pkg/microplus"
)

type Exam struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Exam) Call(ctx context.Context, req *exam.Request, rsp *exam.Response) error {
	log.Info("Received Exam.Call request")
	adminClient := pb.NewAdminService("com.lcb123.srv.admin", microplus.NewClient())
	rsp1, _ := adminClient.Call(ctx, &pb.Request{
		Name: "123456",
	})
	rsp.Msg = "exam_srv " + rsp1.Msg
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Exam) Stream(ctx context.Context, req *exam.StreamingRequest, stream exam.Exam_StreamStream) error {
	log.Infof("Received Exam.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&exam.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Exam) PingPong(ctx context.Context, stream exam.Exam_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&exam.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
