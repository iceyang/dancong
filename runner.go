package dancong

import "log"

type Runner interface {
	PreStart(*Context) interface{}
	Start(*Context) error
	Stop(*Context) error
}

type Runners map[string]Runner

var runners = map[string]Runner{}

func RegisterRunner(name string, runner Runner) {
	if _, ok := runners[name]; ok {
		log.Fatalf("[Dancong] runner %s already exists.", name)
	}
	runners[name] = runner
}

func GetRunner(name string) (runner Runner, ok bool) {
	runner, ok = runners[name]
	return
}
