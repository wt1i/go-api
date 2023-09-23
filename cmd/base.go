package cmd

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/go-god/logger"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type BaseCmd struct {
	cron       *cron.Cron
	rootCtx    context.Context
	rootCancel context.CancelFunc

	cronSchedule string
	resource     string

	beforFunc func() error
	doFunc    func() error
	afterFunc func() error
}

func (bs *BaseCmd) getName() string {
	return bs.resource + "-cmd"
}

func (bs *BaseCmd) Run(ctx context.Context) (err error) {
	if bs.rootCtx != nil {
		return fmt.Errorf("%s still running", bs.getName())
	}

	bs.rootCtx, bs.rootCancel = context.WithCancel(ctx)

	bs.cron = cron.New(cron.WithSeconds())
	if _, err = bs.cron.AddFunc(bs.cronSchedule, func() {
		if err := recover(); err != nil {
			logger.Error(bs.rootCtx, "exec panic error",
				zap.String("module", "web"), zap.String("trace_error", string(debug.Stack())),
			)
		}

		if err := bs.beforFunc(); err != nil {
			logger.Error(bs.rootCtx, "%s 's before error %v", bs.resource, err)
			return
		}

		if err := bs.doFunc(); err != nil {
			logger.Error(bs.rootCtx, "%s 's error %v", bs.resource, err)
			return
		}

		if err := bs.afterFunc(); err != nil {
			logger.Error(ctx, "%s 's after error %v", bs.resource, err)
			return
		}
	}); err != nil {
		return
	}

	bs.cron.Start()
	return
}
