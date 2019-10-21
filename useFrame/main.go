package awesomeProject1

import (
	"awesomeProject1/collector"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)
func init() {
	if err := collector.Init();err!=nil{

log.Println("finish",err.Error())

		panic(err)

	}
}
func main() {

	//监听项目停止信号
	kill := make(chan os.Signal, 1)
	signal.Notify(kill,syscall.SIGINT,syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-kill:
		case <-collector.Errs() :

		}

		cancel()

	}()
	if err := collector.Run(ctx);err!=nil{
		log.Println("finish error --> ", err.Error())



	}else {
		log.Println("finish  --> ")

	}

}
