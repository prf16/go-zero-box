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
		s.MessageMailQueue.Handler(),
		s.MessageSmsQueue.Handler(),
		s.MessageWechatQueue.Handler(),
	}
}
