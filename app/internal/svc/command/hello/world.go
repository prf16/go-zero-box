package hello

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

type World struct {
}

func NewWorld() *World {
	return &World{}
}

// Type 返回任务的唯一标识符
func (s *World) Type() string {
	return "hello:world"
}

// Handle 返回一个 cobra.Command 实例，用于执行脚本任务
func (s *World) Handle() *cobra.Command {
	return &cobra.Command{
		Use:   s.Type(),
		Short: "描述信息，这是一个hello任务",
		Run: func(cmd *cobra.Command, args []string) {
			logx.WithContext(context.Background()).Infof("World World command")
		},
	}
}
