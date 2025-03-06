package message

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-box/app/internal/model/usermodel"
	"go-zero-box/app/internal/services/message"
	"go-zero-box/pkg/asynqx"
)

// 邮箱通知队列

const MailQueueType = "message:mail"

// 队列里的任务消息体

type MailQueuePayload struct {
	User    *usermodel.User
	Content string
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.（编写一个函数NewXXXTask来创建任务。）
// A task consists of a type and a payload.（任务由类型和有效载荷组成。）
//----------------------------------------------

func MailQueueEnqueue(ctx context.Context, payload MailQueuePayload) error {
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	taskInfo, err := asynqx.GetClient().EnqueueContext(ctx, asynq.NewTask(MailQueueType, payloadByte))
	if err != nil {
		return err
	}

	logx.Infof("MailQueueEnqueue taskInfo.ID: %s", taskInfo.ID)
	return nil
}

type MailQueue struct {
	MessageService *message.Service
}

func NewMailQueue(messageService *message.Service) *MailQueue {
	return &MailQueue{MessageService: messageService}
}

func (q *MailQueue) Async() *asynqx.Handler {
	return &asynqx.Handler{
		Type:        MailQueueType,
		Concurrency: 10,
		Async: func(ctx context.Context, t *asynq.Task) error {
			logx.Infof("MailQueue ProcessTask t.Payload: %+v", string(t.Payload()))
			var payload MailQueuePayload
			if err := json.Unmarshal(t.Payload(), &payload); err != nil {
				logx.Errorf("MailQueue ProcessTask json.Unmarshal err: %v", err)
				return err
			}

			err := q.MessageService.Mail(payload.User, payload.Content)
			if err != nil {
				logx.Errorf("MailQueue ProcessTask q.MessageService.MailQueue err: %v payload: %+v", err, payload)
				return err
			}
			return nil
		},
	}
}
