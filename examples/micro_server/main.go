package main

import (
	"github.com/iceyang/dancong"
	"github.com/iceyang/dancong/pkg/runner"
	micro "github.com/micro/go-micro/v2"
)

const serviceId string = "dancong.micro"

func NewGoMicro() micro.Service {
	service := micro.NewService(micro.Name(serviceId))
	service.Init()

	return service
}

func main() {
	opts := dancong.Options(
		dancong.WithBean(
			NewGoMicro,
		),
		dancong.WithRunner(runner.Gomicro2Runner),
	)
	dc := dancong.New(opts)
	dc.Run()
}
