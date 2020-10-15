package dancong

type Runner interface {
	Start(*Context) error
	Stop(*Context) error
}
