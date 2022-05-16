package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/marmotedu/iam/pkg/shutdown"
	"github.com/marmotedu/iam/pkg/shutdown/shutdownmanagers/posixsignal"
)

func main() {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	gs.SetErrorHandler(shutdown.ErrorFunc(func(err error) {
		fmt.Println("Error: ", err)
	}))

	gs.AddShutdownCallback(shutdown.ShutdownFunc(func(shutdownManager string) error {
		fmt.Println("Shutdown callback start")
		time.Sleep(time.Second)
		fmt.Println("Shutdown callback finished")
		return errors.New("myError")
	}))

	if err := gs.Start(); err != nil {
		fmt.Println("Start: ", err)
		return
	}

	q := make(chan struct{})
	<-q
}
