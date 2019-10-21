package collector

import (
	"awesomeProject1/exec"
	"context"
	"sync"
)

var(
	exc exec.Exec

)
var initOnce = sync.Once{}
var errs=make(chan error)
func Init() error {
	var err error
	initOnce.Do(func() {
		err = initialize()
	})
	return err

}
func initialize() error {

	exc = exec.NewExec()

	return nil
}
func Errs() <-chan error{


	return errs
}
func Run(ctx context.Context)error  {

	err:=exc.Run(ctx)
	if err!=nil {
		errs<-err
		return err
	}
	return nil
}
