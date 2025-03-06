package message

import (
	"fmt"
	"go-zero-box/app/internal/model/usermodel"
)

func (s *Service) Sms(user *usermodel.User, content string) error {
	fmt.Printf("发送短信服务 Sms to User: user_id=%d, name=%s\n", user.Id, user.Name)

	return nil
}
