package dancong

import (
	"context"
	"io/ioutil"
	"log"

	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
)

type Option interface {
	apply(*dancong)
}

// WithRunner will apply runner to fx 'invokes',
// Runner will start when app start.
func WithRunner(r Runner) Option {
	return runnerOption{runner: r}
}

type runnerOption struct {
	runner Runner
}

func (o runnerOption) apply(dc *dancong) {
	invokeLifecycle := func(lc fx.Lifecycle) {
		lc.Append(fx.Hook{
			OnStart: func(context.Context) error {
				return o.runner.Start(dc.ctx)
			},
			OnStop: func(context.Context) error {
				return o.runner.Stop(dc.ctx)
			},
		})
	}
	invoke := fx.Invoke(
		o.runner.Init(dc.ctx),
		invokeLifecycle,
	)
	dc.fxOptions = fx.Options(dc.fxOptions, invoke)
}

// WithBean will using constructors as providers
func WithBean(constructors ...interface{}) Option {
	return beanOption{constructors: constructors}
}

type beanOption struct {
	constructors []interface{}
}

func (o beanOption) apply(dc *dancong) {
	dc.fxOptions = fx.Options(
		dc.fxOptions,
		fx.Provide(o.constructors...),
	)
}

type configOption struct {
	filePath string
}

func WithConfig(filePath string) Option {
	return configOption{filePath: filePath}
}

func (o configOption) apply(dc *dancong) {
	path := o.filePath
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Loading config(%s) failed: %s \n", path, err)
	}
	err = yaml.Unmarshal(yamlFile, &dc.ctx.config)
	if err != nil {
		log.Fatalf("Loading config(%s) failed: %s \n", path, err)
	}
}

// Option Group
type options []Option

func Options(opts ...Option) Option {
	return options(opts)
}

func (opts options) apply(dc *dancong) {
	for _, opt := range opts {
		opt.apply(dc)
	}
}
