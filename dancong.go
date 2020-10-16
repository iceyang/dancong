package dancong

import (
	"go.uber.org/fx"
)

type dancong struct {
	// Dancong Context
	ctx *Context
	// the initial options
	opts options

	// fx module
	fxApp     *fx.App
	fxOptions fx.Option
}

// Create a dancong application
func New(opts ...Option) *dancong {
	dc := &dancong{
		ctx:       &Context{},
		opts:      opts,
		fxOptions: fx.Options(),
	}

	for _, opt := range opts {
		opt.apply(dc)
	}

	dc.fxApp = fx.New(dc.fxOptions)
	return dc
}

// Start Application
func (dc *dancong) Run() {
	dc.fxApp.Run()
}
