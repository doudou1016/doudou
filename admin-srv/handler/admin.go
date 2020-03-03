package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	admin "doudou/admin-srv/proto/admin"
)

type Admin struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Admin) Call(ctx context.Context, req *admin.Request, rsp *admin.Response) error {
	log.Info("Received Admin.Call request")
	rsp.Msg = "admin_srv " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Admin) Stream(ctx context.Context, req *admin.StreamingRequest, stream admin.Admin_StreamStream) error {
	log.Infof("Received Admin.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&admin.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Admin) PingPong(ctx context.Context, stream admin.Admin_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&admin.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
