package retry

import (
	"github.com/avast/retry-go"
)

type Runner struct {
	opts []retry.Option
}

func NewRunner(opts []retry.Option) *Runner {
	return &Runner{opts: opts}
}

func (r *Runner) Do(fn func() error) error {
	return retry.Do(
		fn,
		r.opts...,
	)
}
