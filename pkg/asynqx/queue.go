package asynqx

import (
	"log"

	"github.com/hibiken/asynq"
)

// NewClient 生产者客户端
func NewClient(c *Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
}

// Server 消费者
type Server struct {
	server  *asynq.Server
	handler *Handler
}

func NewServer(config *Config, handler *Handler) *Server {
	if handler.Concurrency == 0 {
		handler.Concurrency = 1
	}
	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.DB,
		},
		asynq.Config{
			Concurrency: handler.Concurrency,
		},
	)

	return &Server{
		server:  server,
		handler: handler,
	}
}

func (q *Server) Start() {
	mux := asynq.NewServeMux()
	mux.Handle(q.handler.Type, q.handler.Handler)

	if err := q.server.Start(mux); err != nil {
		panic(err)
	}

	log.Printf("[server:queue] start success at Type: %s Concurrency: %d...", q.handler.Type, q.handler.Concurrency)
}

func (q *Server) Stop() {
	q.server.Stop()
	log.Printf("[server:queue] stop Type: %s", q.handler.Type)
}
