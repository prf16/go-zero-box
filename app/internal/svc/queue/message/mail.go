package message

import (
	"context"
	"encoding/json"
	"go-zero-box/app/internal/svc/model/usermodel"
	"go-zero-box/app/internal/svc/services/message"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

// 邮箱通知队列

const MailQueueType = "message:mail"

// 队列里的任务消息体

type MailQueuePayload struct {
	User    *usermodel.User
	Content string
}

// ----------------------------------------------
// XXXQueueEnqueue 入队函数
// Write a function NewXXXTask to create a task.（编写一个函数NewXXXTask来创建任务。）
// A task consists of a type and a payload.（任务由类型和有效载荷组成。）
// ----------------------------------------------

func MailQueueEnqueue(ctx context.Context, client *asynq.Client, payload MailQueuePayload) error {
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		logx.Errorf("MailQueueEnqueue json.Marshal err: %v", err)
		return nil
	}

	taskInfo, err := client.EnqueueContext(ctx, asynq.NewTask(MailQueueType, payloadByte))
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

// ProcessTask 处理任务函数
func (q *MailQueue) ProcessTask(ctx context.Context, t *asynq.Task) error {
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
}
