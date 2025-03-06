package message

import (
	"fmt"
	"go-zero-box/app/internal/model/usermodel"
)

func (s *Service) Mail(user *usermodel.User, content string) error {
	fmt.Printf("发送邮箱服务 Email to User: user_id=%d, name=%s\n", user.Id, user.Name)

	return nil
}
