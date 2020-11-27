package dancong

import (
	"context"
	"io/ioutil"

	logrus "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
)

type Option interface {
	apply(*Dancong)
}

// WithRunner will apply runner to fx 'invokes',
// Runner will start when app start.
func WithRunner(runnerNames ...string) Option {
	return runnerOption{runnerNames: runnerNames}
}

type runnerOption struct {
	runnerNames []string
}

func (o runnerOption) apply(dc *Dancong) {
	for _, name := range o.runnerNames {
		runner, ok := GetRunner(name)
		if !ok {
			dc.logger.Fatalf("[Dancong] runner %s is not exists.", name)
		}
		invokeLifecycle := func(lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: func(context.Context) error {
					return runner.Start(dc)
				},
				OnStop: func(context.Context) error {
					return runner.Stop(dc)
				},
			})
		}
		invoke := fx.Invoke(
			runner.PreStart(dc),
			invokeLifecycle,
		)
		dc.fxOptions = fx.Options(dc.fxOptions, invoke)
	}
}

// WithBean will using constructors as providers
func WithBean(constructors ...interface{}) Option {
	return beanOption{constructors: constructors}
}

type beanOption struct {
	constructors []interface{}
}

func (o beanOption) apply(dc *Dancong) {
	dc.fxOptions = fx.Options(
		dc.fxOptions,
		fx.Provide(o.constructors...),
	)
}

func WithLogger(logger *logrus.Logger) Option {
	return loggerOption{logger: logger}
}

type loggerOption struct {
	logger *logrus.Logger
}

func (o loggerOption) apply(dc *Dancong) {
	dc.logger = o.logger
	dc.fxOptions = fx.Options(
		dc.fxOptions,
		fx.Logger(o.logger),
	)
}

type configOption struct {
	filePath string
}

func WithConfig(filePath string) Option {
	return configOption{filePath: filePath}
}

func (o configOption) apply(dc *Dancong) {
	path := o.filePath
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		dc.logger.Fatalf("Loading config(%s) failed: %s \n", path, err)
	}
	var config map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		dc.logger.Fatalf("Loading config(%s) failed: %s \n", path, err)
	}
	dc.ctx.SetConfig(config)
}

// Option Group
type options []Option

func Options(opts ...Option) Option {
	return options(opts)
}

func (opts options) apply(dc *Dancong) {
	for _, opt := range opts {
		opt.apply(dc)
	}
}
