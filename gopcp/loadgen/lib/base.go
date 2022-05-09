package lib

import "time"

type RetCode int

const (
	RET_CODE_SUCCESS              RetCode = 0
	RET_CODE_WARNING_CALL_TIMEOUT RetCode = 1001
	RET_CODE_ERROR_CALL           RetCode = 2001
	RET_CODE_ERROR_RESPONSE       RetCode = 2002
	RET_CODE_ERROR_CALEE          RetCode = 2003
	RET_CODE_FATAL_CALL           RetCode = 3001
)

const (
	STATUS_ORIGINAL uint32 = 0
	STATUS_STARTING uint32 = 1
	STATUS_STARTED  uint32 = 2
	STATUS_STOPPING uint32 = 3
	STATUS_STOPPED  uint32 = 4
)

type RawReq struct {
	ID  int64
	Req []byte
}

type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

type CallResult struct {
	ID     int64
	Req    RawReq
	Resp   RawResp
	Code   RetCode
	Msg    string
	Elapse time.Duration
}

type Generator interface {
	Start() bool
	Stop() bool
	Status() uint32
	CallCount() int64
}
