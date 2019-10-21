package exec

import (
	"context"
	"github.com/pkg/errors"
)

type Exec interface {
	Run(ctx context.Context)error
}

type exec struct {}

func  NewExec() Exec {
	return &exec{}
}
func (e *exec)Run(ctx context.Context)  error{
	return e.run(ctx)
}

func (e *exec) run(ctx context.Context) error {

	return errors.New("analyze_statistical程序退出")
}
