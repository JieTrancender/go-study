package loadgen

import (
	"time"

	"lib"
)

type ParamSet struct {
	Caller     lib.Caller
	TimeoutNS  time.Duration
	LPS        uint32
	DurationNS time.Duration
	ResultCh   chan *lib.CallResult
}
