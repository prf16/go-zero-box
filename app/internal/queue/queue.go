package queue

import (
	"github.com/google/wire"
	"go-zero-box/app/internal/queue/message"
	"go-zero-box/pkg/asynqx"
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

func Handler(s *Queue) []*asynqx.Handler {
	return []*asynqx.Handler{
		s.MessageMailQueue.Async(),
		s.MessageSmsQueue.Async(),
		s.MessageWechatQueue.Async(),
	}
}
