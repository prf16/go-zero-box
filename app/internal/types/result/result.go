package result

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"runtime/debug"
)

type Result struct {
	Code    any    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (e *Result) Error() string {
	return e.Message
}

const CodeDefault = 10001 // 默认异常错误码，前端可提示
const CodeSystem = 20001  // 系统异常错误码，前端可不提示
const CodeAuth = 40000    // 未登录状态码

func ResponseAuth(ctx context.Context, msg string) error {
	logx.WithContext(ctx).Errorf("异常信息：%s 堆栈追踪：%s", msg, debug.Stack())
	return &Result{Code: CodeAuth, Message: msg}
}

func Response(ctx context.Context, msg string) error {
	logx.WithContext(ctx).Errorf("异常信息：%s 堆栈追踪：%s", msg, debug.Stack())
	return &Result{Code: CodeDefault, Message: msg}
}

func ResponseSystem(ctx context.Context, msg string) error {
	logx.WithContext(ctx).Errorf("异常信息：%s 堆栈追踪：%s", msg, debug.Stack())
	return &Result{Code: CodeSystem, Message: msg}
}
