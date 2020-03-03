package microplus

import (
	"time"

	"github.com/micro/go-micro/v2/client"
	grpcC "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/service"
	grpcS "github.com/micro/go-micro/v2/service/grpc"
)

func etcdOptions(op *registry.Options) {
	op.Addrs = []string{
		"http://127.0.0.1:2379",
	}
	op.Timeout = 5 * time.Second
}

// NewService returns a new go-micro service
func NewService(opts ...service.Option) service.Service {
	//etcd registry
	reg := etcd.NewRegistry(etcdOptions)

	// set the registry and selector
	options := []service.Option{
		service.Registry(reg),
		service.RegisterTTL(time.Second * 15),      //重新注册时间
		service.RegisterInterval(time.Second * 10), //注册过期时间
		service.Version("latest"),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return grpcS.NewService(options...)
}

// NewClient returns a new go-micro clinet
func NewClient(opts ...client.Option) client.Client {
	//etcd registry
	r := etcd.NewRegistry(etcdOptions)

	// set the registry and selector
	options := []client.Option{
		client.Registry(r),
	}
	// append user options
	options = append(options, opts...)
	// return a micro.Client
	return grpcC.NewClient(options...)
}
