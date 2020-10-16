package runner

import (
	"fmt"
	"log"
	"net"

	"github.com/iceyang/dancong"
	"google.golang.org/grpc"
)

type GrpcRunner struct {
	BaseRunner

	Server *grpc.Server

	// listening address
	Addr string
}

func (runner *GrpcRunner) Start(ctx *dancong.Context) error {
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(runner.Addr))
		if err != nil {
			log.Fatalf("listen: %s\n", err)
		}
		runner.Server.Serve(lis)
	}()
	return nil
}

func (runner *GrpcRunner) Stop(ctx *dancong.Context) error {
	return nil
}
