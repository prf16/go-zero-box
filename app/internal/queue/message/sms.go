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

const SmsQueueType = "message:sms"

// 队列里的任务消息体

type SmsQueuePayload struct {
	User    *usermodel.User
	Content string
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.（编写一个函数NewXXXTask来创建任务。）
// A task consists of a type and a payload.（任务由类型和有效载荷组成。）
//----------------------------------------------

func SmsQueueEnqueue(ctx context.Context, payload SmsQueuePayload) error {
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	taskInfo, err := asynqx.GetClient().EnqueueContext(ctx, asynqx.NewTask(SmsQueueType, payloadByte))
	if err != nil {
		return err
	}

	logx.Infof("SmsQueueEnqueue taskInfo: %+v", taskInfo)
	return nil
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.（编写一个函数 HandleXXXTask 来处理输入的任务）
// Note that it satisfies the asynq.HandlerFunc interface.（请注意它满足 asynq.HandlerFunc 接口。）
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.（处理程序不一定需要是一个函数。你可以定义一个满足 asynq.Handler 接口的类型。请参考下面的示例。）
//---------------------------------------------------------------

type SmsQueue struct {
	MessageService *message.Service
}

func NewSmsQueue(messageService *message.Service) *SmsQueue {
	return &SmsQueue{MessageService: messageService}
}

// Async 异步处理任务
func (q *SmsQueue) Async() *asynqx.Handler {
	return &asynqx.Handler{
		Type:        SmsQueueType,
		Concurrency: 10,
		Async: func(ctx context.Context, t *asynq.Task) error {
			logx.Infof("SmsQueue ProcessTask t.Payload: %+v", string(t.Payload()))
			var payload SmsQueuePayload
			if err := json.Unmarshal(t.Payload(), &payload); err != nil {
				logx.Errorf("SmsQueue ProcessTask json.Unmarshal err: %v", err)
				return err
			}

			err := q.MessageService.Sms(payload.User, payload.Content)
			if err != nil {
				logx.Errorf("SmsQueue ProcessTask q.MessageService.SmsQueue err: %v payload: %+v", err, payload)
				return err
			}
			return nil
		},
	}
}
