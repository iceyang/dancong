package dancong

type Runner interface {
	Init(*Context) interface{}
	Start(*Context) error
	Stop(*Context) error
}

type Runners map[string]Runner

var runners = map[string]Runner{}

func RegisterRunner(name string, runner Runner) {
	runners[name] = runner
}

func GetRunner(name string) (runner Runner, ok bool) {
	runner, ok = runners[name]
	return
}
