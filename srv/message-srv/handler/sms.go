package handler

import (
	"context"
	"github.com/iron-kit/go-ironic/micro-assistant"
	sms "iunite.club/srv/message-srv/proto/sms"
)

type SMSHandler struct {
	assistant.BaseHandler
}

func (s *SMSHandler) SendVerifyCode(ctx context.Context, in *sms.SendVerifyCodeRequest, resp *sms.SendResponse) error {

	return nil
}

func (s *SMSHandler) ValidateMobileCode(ctx context.Context, in *sms.ValidateMobileCodeRequest, resp *sms.SendResponse) error {
	return nil
}
