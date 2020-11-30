package dancong

import (
	"github.com/iceyang/dancong/internal/log"
	"go.uber.org/fx"
)

type Dancong struct {
	// Dancong Context
	ctx *Context
	// the initial options
	opts options

	logger log.Printer

	// fx module
	fxApp     *fx.App
	fxOptions fx.Option
}

// Create a Dancong application
func New(opts ...Option) *Dancong {
	logger := log.DefaultLogger()

	dc := &Dancong{
		ctx:       NewContext(),
		opts:      opts,
		logger:    logger,
		fxOptions: fx.Logger(logger),
	}

	for _, opt := range opts {
		opt.apply(dc)
	}

	dc.fxApp = fx.New(dc.fxOptions)
	return dc
}

func (dc *Dancong) GetContext() *Context {
	return dc.ctx
}

func (dc *Dancong) GetLogger() log.Printer {
	return dc.logger
}

// Start Application
func (dc *Dancong) Run() {
	dc.fxApp.Run()
}
