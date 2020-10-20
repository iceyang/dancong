package runner

import (
	"log"

	"github.com/iceyang/dancong"
	micro "github.com/micro/go-micro/v2"
)

const Gomicro2Runner = "gomicro2Runner"

type gomicro2Runner struct {
	BaseRunner

	service micro.Service
}

func init() {
	dancong.RegisterRunner(Gomicro2Runner, &gomicro2Runner{})
}

func (runner *gomicro2Runner) PreStart(ctx *dancong.Context) interface{} {
	return func(service micro.Service) {
		runner.service = service
	}
}

func (runner *gomicro2Runner) Start(ctx *dancong.Context) error {
	log.Println("Starting GoMicro server.")
	go func() {
		if err := runner.service.Run(); err != nil {
			log.Fatalf("gomicro listen: %s\n", err)
		}
	}()
	return nil
}
