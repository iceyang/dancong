package runner

import (
	"github.com/iceyang/dancong"
)

type BaseRunner struct{}

func (runner *BaseRunner) Init(ctx *dancong.Context) interface{} {
	return func() {}
}

func (runner *BaseRunner) Start(ctx *dancong.Context) error {
	return nil
}

func (runner *BaseRunner) Stop(ctx *dancong.Context) error {
	return nil
}
