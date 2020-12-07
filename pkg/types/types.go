package types

import (
	"context"
)

var (
	defaultCtx context.Context
)

const (
	// WorkerNum is the number of worker goroutines.
	WorkerNum int = 2
)

func init() {
	defaultCtx = context.Background()
}

func GetDefaultCtx() context.Context {
	return defaultCtx
}
