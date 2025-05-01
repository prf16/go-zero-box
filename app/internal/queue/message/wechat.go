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

// 任务的唯一标识符

const WechatQueueType = "message:wechat"

// 队列里的任务消息体

type WechatQueuePayload struct {
	User    *usermodel.User
	Content string
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.（编写一个函数NewXXXTask来创建任务。）
// A task consists of a type and a payload.（任务由类型和有效载荷组成。）
//----------------------------------------------

func WechatQueueEnqueue(ctx context.Context, payload WechatQueuePayload) error {
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	taskInfo, err := asynqx.Client.EnqueueContext(ctx, asynqx.NewTask(WechatQueueType, payloadByte))
	if err != nil {
		return err
	}

	logx.Infof("Queue MailProcess EnqueueWechatQueue taskInfo: %+v", taskInfo)
	return nil
}

type WechatQueue struct {
	MessageService *message.Service
}

func NewWechatQueue(messageService *message.Service) *WechatQueue {
	return &WechatQueue{MessageService: messageService}
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.（编写一个函数 HandleXXXTask 来处理输入的任务）
// Note that it satisfies the asynq.HandlerFunc interface.（请注意它满足 asynq.HandlerFunc 接口。）
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.（处理程序不一定需要是一个函数。你可以定义一个满足 asynq.Handler 接口的类型。请参考下面的示例。）
//---------------------------------------------------------------

// Async 异步处理任务
func (q *WechatQueue) Async() *asynqx.Handler {
	return &asynqx.Handler{
		Type:        SmsQueueType,
		Concurrency: 10,
		Async: func(ctx context.Context, t *asynq.Task) error {
			logx.Infof("WechatQueue ProcessTask t.Payload: %s", string(t.Payload()))
			var payload WechatQueuePayload
			if err := json.Unmarshal(t.Payload(), &payload); err != nil {
				logx.Errorf("WechatQueue ProcessTask json.Unmarshal err: %v", err)
				return err
			}

			err := q.MessageService.Wechat(payload.User, payload.Content)
			if err != nil {
				logx.Errorf("WechatQueue ProcessTask q.MessageService.Wechat err: %v payload: %+v", err, payload)
				return err
			}
			return nil
		},
	}
}
