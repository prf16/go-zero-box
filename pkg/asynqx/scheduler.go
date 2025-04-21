package asynqx

import (
	"fmt"
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
		log.Fatalf("| Server: scheduler | load location error: %v", err)
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
			panic(fmt.Sprintf("| Server: scheduler | q.scheduler.Register error: %v", err))
		} else {
			log.Printf("| Server: scheduler | type: %s | Register at %s %s", v.Type, entryID, v.Scheduler)
		}
	}

	log.Printf("| Server: scheduler | run...")
	if err := q.scheduler.Start(); err != nil {
		panic(fmt.Sprintf("| Server: scheduler | q.scheduler.Run error: %v", err))
	}

	select {}
}
func (q *Scheduler) Stop() {
	q.scheduler.Shutdown()
	log.Printf("| Server: scheduler | stop")
}
