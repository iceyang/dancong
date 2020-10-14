package runner

import "github.com/iceyang/dancong"

type Runner interface {
	Start(*dancong.Context) error
	Stop(*dancong.Context) error
}
