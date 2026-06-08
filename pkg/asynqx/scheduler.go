package asynqx

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
)

type Scheduler struct {
	scheduler *asynq.Scheduler
	server    []*Server
}

func NewScheduler(config *Config, handler []*Handler) *Scheduler {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	var (
		scheduler = asynq.NewScheduler(
			&asynq.RedisClientOpt{
				Addr:     config.Addr,
				Password: config.Password,
				DB:       config.DB,
			},
			&asynq.SchedulerOpts{
				Location: loc,
			},
		)

		server []*Server
	)

	for _, v := range handler {
		entryID, err := scheduler.Register(v.Scheduler, asynq.NewTask(v.Type, nil))
		if err != nil {
			panic(err)
		}

		server = append(server, NewServer(config, v))

		log.Printf("[server:scheduler] register Type: %s entryID: %s Scheduler: %s", v.Type, entryID, v.Scheduler)
	}

	return &Scheduler{
		scheduler: scheduler,
		server:    server,
	}
}

func (q *Scheduler) Start() {
	for _, v := range q.server {
		v.Start()
	}

	log.Printf("[server:scheduler] run...")
	if err := q.scheduler.Run(); err != nil {
		panic(err)
	}
}

func (q *Scheduler) Stop() {
	for _, v := range q.server {
		v.Stop()
	}

	q.scheduler.Shutdown()
	log.Printf("[server:scheduler] stop")
}
