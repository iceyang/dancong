package runner

import (
	"fmt"
	"net"

	"github.com/iceyang/dancong"
	"google.golang.org/grpc"
)

const GrpcRunner = "grpcRunner"

type grpcRunner struct {
	BaseRunner

	server *grpc.Server
}

func init() {
	dancong.RegisterRunner(GrpcRunner, &grpcRunner{})
}

func (runner *grpcRunner) PreStart(dc *dancong.Dancong) interface{} {
	return func(server *grpc.Server) {
		runner.server = server
	}
}

func (runner *grpcRunner) Start(dc *dancong.Dancong) error {
	ctx := dc.GetContext()
	v, _ := ctx.GetConfig("grpc.addr")
	addr := v.(string)
	dc.GetLogger().Infof("Starting GRPC server. Listening at %s\n", addr)
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(addr))
		if err != nil {
			dc.GetLogger().Fatalf("listen: %s\n", err)
		}
		runner.server.Serve(lis)
	}()
	return nil
}
