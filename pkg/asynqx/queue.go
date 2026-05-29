package asynqx

import (
	"go-zero-box/pkg/redis"
	"log"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/service"
)

type Queue struct {
	server  *asynq.Server
	handler *Handler
}

func NewQueue(config *redis.Config, handler *Handler) service.Service {
	concurrency := handler.Concurrency
	if concurrency == 0 {
		concurrency = 1
	}
	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     config.Host,
			Password: config.Pass,
		},
		asynq.Config{
			Concurrency: concurrency,
		},
	)

	return &Queue{
		server:  server,
		handler: handler,
	}
}

func (q *Queue) Start() {
	mux := asynq.NewServeMux()
	mux.HandleFunc(q.handler.Type, q.handler.Handler)

	if err := q.server.Start(mux); err != nil {
		panic(err)
	}

	log.Printf("[server:queue] start success at Type: %s Concurrency: %d...", q.handler.Type, q.handler.Concurrency)
}
func (q *Queue) Stop() {
	q.server.Stop()
	log.Printf("[server:queue] stop Type: %s", q.handler.Type)
}
