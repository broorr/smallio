package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SyncOperation(f func() error) {
	oneSigChan := make(chan os.Signal, 1)
	// 申请信号量锁
	signal.Notify(oneSigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		err := f()
		if err != nil {
			panic(fmt.Sprintf("SyncOperation ==> %v", err))
		}
	}()
	sig := <-oneSigChan
	log.Printf("Caught SIGTERM %v", sig)
}
