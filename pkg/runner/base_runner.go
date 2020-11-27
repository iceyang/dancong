package runner

import (
	"github.com/iceyang/dancong"
)

type BaseRunner struct{}

func (runner *BaseRunner) PreStart(dc *dancong.Dancong) interface{} {
	return func() {}
}

func (runner *BaseRunner) Start(dc *dancong.Dancong) error {
	return nil
}

func (runner *BaseRunner) Stop(dc *dancong.Dancong) error {
	return nil
}
