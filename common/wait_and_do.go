package common

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/go-god/logger"
	"go.uber.org/zap"
)

// WaitAndDo 每间隔 interval 时间执行一次 doFunc
func WaitAndDo(ctx context.Context, interval time.Duration, doFunc func() error) {
	tick := time.NewTicker(interval)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			if err := doFunc(); err != nil {
				logger.Error(ctx, "exec do func error",
					zap.String("module", "web"), zap.String("trace_error", string(debug.Stack())),
				)
			}
		case <-ctx.Done():
			return
		}
	}
}
