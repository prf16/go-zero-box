package message

import (
	"fmt"
	"go-zero-box/app/internal/model/usermodel"
)

func (s *Service) Wechat(user *usermodel.User, content string) error {
	fmt.Printf("发送微信通知服务 Wechat to User: user_id=%d, name=%s\n", user.Id, user.Name)

	return nil
}
