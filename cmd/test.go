package cmd

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type TestTask struct {
	BaseCmd
}

func NewTestTask() *TestTask {
	return &TestTask{
		BaseCmd: BaseCmd{
			cron:    &cron.Cron{},
			rootCtx: nil,
			rootCancel: func() {
			},
			cronSchedule: fmt.Sprintf("@every %ds", 5),
			resource:     "test",
			beforFunc: func() error {
				fmt.Println("before do")
				return nil
			},
			doFunc: func() error {
				fmt.Println("do")
				return nil
			},
			afterFunc: func() error {
				fmt.Println("after do")
				return nil
			},
		},
	}
}
