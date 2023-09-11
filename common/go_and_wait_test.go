package common

import (
	"context"
	"errors"
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
					func(context.Context) error {
						return errors.New("1")
					},
					func(context.Context) error {
						return errors.New("2")
					},
					func(context.Context) error {
						return errors.New("3")
					},
					func(context.Context) error {
						return errors.New("4")
					},
					func(context.Context) error {
						panic("panic")
					},
					func(context.Context) error {
						return nil
					},
					func(context.Context) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errList := GoAndWait(context.Background(), tt.args.handlers)

			if len(errList.FilterNil()) != 5 {
				fmt.Println(len(errList))
				t.Error("GoAndWait get error has err")
			}
		})
	}
}
