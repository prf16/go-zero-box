package queue

import (
	"go-zero-box/app/internal/svc/queue/message"
	"go-zero-box/pkg/asynqx"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewQueue,
	message.NewMailQueue,
	message.NewSmsQueue,
	message.NewWechatQueue,
)

type Queue struct {
	MessageMailQueue   *message.MailQueue
	MessageSmsQueue    *message.SmsQueue
	MessageWechatQueue *message.WechatQueue
}

func NewQueue(messageMailQueue *message.MailQueue, messageSmsQueue *message.SmsQueue, messageWechatQueue *message.WechatQueue) *Queue {
	return &Queue{MessageMailQueue: messageMailQueue, MessageSmsQueue: messageSmsQueue, MessageWechatQueue: messageWechatQueue}
}

func Register(s *Queue) []*asynqx.Handler {
	return []*asynqx.Handler{
		{Type: message.MailQueueType, Concurrency: 10, Handler: s.MessageMailQueue},
		{Type: message.SmsQueueType, Concurrency: 10, Handler: s.MessageSmsQueue},
		{Type: message.WechatQueueType, Concurrency: 10, Handler: s.MessageWechatQueue},
	}
}
