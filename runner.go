package dancong

type Runner interface {
	Init(*Context) interface{}
	Start(*Context) error
	Stop(*Context) error
}
