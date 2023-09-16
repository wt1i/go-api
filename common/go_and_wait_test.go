package common

import (
	"context"
	"fmt"
	"testing"
)

func TestGoAndWait(t *testing.T) {
	type args struct {
		handlers []ContextHandle
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				[]ContextHandle{
					func(ctx context.Context, i int) error {
						return fmt.Errorf("%v", i)
					},
					func(ctx context.Context, i int) error {
						return fmt.Errorf("%v", i)
					},
					func(ctx context.Context, i int) error {
						return fmt.Errorf("%v", i)
					},
					func(ctx context.Context, i int) error {
						return fmt.Errorf("%v", i)
					},
					func(context.Context, int) error {
						panic("panic")
					},
					func(context.Context, int) error {
						return nil
					},
					func(context.Context, int) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errList := GoAndWait(context.Background(), tt.args.handlers)
			fmt.Println(errList)
			if len(errList.FilterNil()) != 5 {
				t.Error("GoAndWait get error has err")
			}
		})
	}
}
