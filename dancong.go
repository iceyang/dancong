package dancong

import (
	"context"

	"github.com/iceyang/dancong/pkg/runner"
	"go.uber.org/fx"
)

type dancong struct {
	app fx.App

	// Dancong Context
	ctx *Context

	// fxOptions
	fxOptions fx.Option
}

func New() *dancong {
	return &dancong{
		fxOptions: fx.Options(),
	}
}

// Invoke Runner
func (dc *dancong) AddRunner(r runner.Runner) {
	dc.fxOptions = fx.Options(
		dc.fxOptions,
		fx.Invoke(func(lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: func(context.Context) error {
					return r.Start(dc.ctx)
				},
				OnStop: func(context.Context) error {
					return r.Stop(dc.ctx)
				},
			})
		}),
	)
}

// Start Application
func (dc *dancong) Run() {
	dc.app = fx.New(dc.fxOptions)
	dc.app.Run()
}
