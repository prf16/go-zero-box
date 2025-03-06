package demo

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-box/pkg/asynqx"
	"time"
)

type PersonProcess struct {
}

func NewPersonProcess() *PersonProcess {
	return &PersonProcess{}
}

// Type 返回任务的唯一标识符
func (s *PersonProcess) Type() string {
	return "demo:person"
}

// Sync 返回一个 cobra.Command 实例，用于同步执行任务
func (s *PersonProcess) Sync() *cobra.Command {
	return &cobra.Command{
		Use:   s.Type(),
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			s.Handle()
		},
	}
}

// Async 返回一个 asynqx.Handler 实例，用于异步处理任务
func (s *PersonProcess) Async() *asynqx.Handler {
	return &asynqx.Handler{
		Type:      s.Type(),
		Scheduler: "* * * * *",
		Async: func(ctx context.Context, t *asynq.Task) error {
			s.Handle()
			return nil
		},
	}
}
func (s *PersonProcess) Handle() {
	logx.WithContext(context.Background()).Infof("time: %s I'm lisa\n", time.Now().String())
}
