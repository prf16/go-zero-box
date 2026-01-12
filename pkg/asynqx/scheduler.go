package asynqx

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/service"
	"log"
	"time"
)

type Scheduler struct {
	config    *Config
	scheduler *asynq.Scheduler
	handler   []*Handler
}

func NewScheduler(config *Config, handler []*Handler) service.Service {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	return &Scheduler{
		config: config,
		scheduler: asynq.NewScheduler(
			&asynq.RedisClientOpt{
				Addr:     config.Addr,
				Password: config.Password,
				DB:       config.DB,
			},
			&asynq.SchedulerOpts{
				Location: loc,
			},
		),
		handler: handler,
	}
}

func (q *Scheduler) Start() {
	for _, v := range q.handler {
		entryID, err := q.scheduler.Register(v.Scheduler, asynq.NewTask(v.Type, nil))
		if err != nil {
			panic(err)
		}

		log.Printf("[server:scheduler] register Type: %s entryID: %s Scheduler: %s", v.Type, entryID, v.Scheduler)
	}

	if err := q.scheduler.Start(); err != nil {
		panic(err)
	}

	log.Printf("[server:scheduler] run...")
	select {}
}
func (q *Scheduler) Stop() {
	q.scheduler.Shutdown()
	log.Printf("[server:scheduler] Shutdown")
}
