package asynqx

import (
	"log"

	"go-zero-box/pkg/redis"

	"github.com/hibiken/asynq"
)

// NewClient 生产者客户端
func NewClient(c *redis.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.Host,
		Password: c.Pass,
	})
}

// Server 消费者
type Server struct {
	server  *asynq.Server
	handler *Handler
}

func NewServer(config *redis.Config, handler *Handler) *Server {
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

	return &Server{
		server:  server,
		handler: handler,
	}
}

func (q *Server) Start() {
	mux := asynq.NewServeMux()
	mux.HandleFunc(q.handler.Type, q.handler.Handler)

	if err := q.server.Start(mux); err != nil {
		panic(err)
	}

	log.Printf("[server:queue] start success at Type: %s Concurrency: %d...", q.handler.Type, q.handler.Concurrency)
}

func (q *Server) Stop() {
	q.server.Stop()
	log.Printf("[server:queue] stop Type: %s", q.handler.Type)
}
