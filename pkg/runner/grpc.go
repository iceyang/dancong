package runner

import (
	"fmt"
	"log"
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

func (runner *grpcRunner) PreStart(ctx *dancong.Context) interface{} {
	return func(server *grpc.Server) {
		runner.server = server
	}
}

func (runner *grpcRunner) Start(ctx *dancong.Context) error {
	v, _ := ctx.GetConfig("grpc.addr")
	addr := v.(string)
	log.Printf("[Dancong] Starting GRPC server. Listening at %s\n", addr)
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(addr))
		if err != nil {
			log.Fatalf("listen: %s\n", err)
		}
		runner.server.Serve(lis)
	}()
	return nil
}
