package loadgen

import (
	"context"
	"log"
	"time"

	"lib"
)

type myGenerator struct {
	caller      lib.Caller
	timeoutNS   time.Duration
	lps         uint32
	durationNS  time.Duration
	concurrency uint32
	tickets     lib.GoTickets
	ctx         context.Context
	cancelFunc  context.CancelFunc
	callCount   int64
	status      uint32
	resultCh    chan *lib.CallResult
}

func NewGenerator(pset ParamSet) (lib.Generator, error) {
	log.Println("hello world")
	return nil, nil
}
