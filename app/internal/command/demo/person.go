package demo

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-box/pkg/asynqx"
	"time"
)

type Person struct {
}

func NewPerson() *Person {
	return &Person{}
}

// Type 返回任务的唯一标识符
func (s *Person) Type() string {
	return "demo:person"
}

// Sync 返回一个 cobra.Command 实例，用于执行脚本任务
func (s *Person) Sync() *cobra.Command {
	return &cobra.Command{
		Use:   s.Type(),
		Short: "描述信息，这是一个demo任务",
		Run: func(cmd *cobra.Command, args []string) {
			s.Handle()
		},
	}
}

// Async 返回一个 asynqx.Handler 实例，用于执行计划任务
func (s *Person) Async() *asynqx.Handler {
	return &asynqx.Handler{
		Type:      s.Type(),
		Scheduler: "* * * * *",
		Async: func(ctx context.Context, t *asynq.Task) error {
			s.Handle()
			return nil
		},
	}
}
func (s *Person) Handle() {
	logx.WithContext(context.Background()).Infof("time: %s I'm lisa\n", time.Now().String())
}
