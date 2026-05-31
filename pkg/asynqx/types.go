package asynqx

import (
	"context"

	"github.com/hibiken/asynq"
)

type Handler struct {
	Type        string
	Scheduler   string // schedule | 计划任务的调度程序表达式
	Concurrency int    // queue | 队列的并发数
	Handler     func(ctx context.Context, task *asynq.Task) error
}
