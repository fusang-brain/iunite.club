package handler

import (
	"context"
	"encoding/json"
	"github.com/iron-kit/go-ironic"
	go_api "github.com/micro/go-api/proto"
	sms "iunite.club/services/message/proto/sms"
	"iunite.club/services/navo/client"
	"strconv"
)

type SMSHandler struct {
	ironic.BaseHandler
}

func (s *SMSHandler) SendVerifyCode(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	smsService, ok := client.SMSServiceFromContext(ctx)

	if !ok {
		return s.Error(ctx).InternalServerError("SMS service is not exists")
	}

	params := struct {
		Mobile string `json:"mobile,omitempty" validate:"nonzero"`
		Type   string `json:"type,omitempty" validate:"nonzero"`
	}{}

	if err := json.Unmarshal([]byte(req.Body), &params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).InternalServerError(err.Error()))
	}

	if err := s.Validate(&params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).BadRequest(err.Error()))
	}

	if smsResp, err := smsService.SendVerifyCode(ctx, &sms.SendVerifyCodeRequest{
		Mobile: params.Mobile,
		Type:   params.Type,
	}); err != nil {
		// rsp.Body = APIError(err.Error()).String()
		return ErrorResponse(rsp, err)
	} else {
		// rsp.Body = smsResp.
		rsp.Body = APISuccess(map[string]interface{}{
			"msgBody": smsResp.Message,
		}).String()
	}

	return nil
}

func (s *SMSHandler) ValidateSimpleCode(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	smsService, ok := client.SMSServiceFromContext(ctx)

	if !ok {
		return s.Error(ctx).InternalServerError("SMS service is not exists")
	}

	params := struct {
		Mobile string `json:"mobile,omitempty" validate:"nonzero"`
		Code   int    `json:"code,omitempty" validate:"nonzero"`
	}{}

	if err := json.Unmarshal([]byte(req.Body), &params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).InternalServerError(err.Error()))
	}

	if err := s.Validate(&params); err != nil {
		return ErrorResponse(rsp, s.Error(ctx).BadRequest(err.Error()))
	}

	if smsResp, err := smsService.ValidateMobileCode(ctx, &sms.ValidateMobileCodeRequest{
		Mobile: params.Mobile,
		Code:   strconv.Itoa(params.Code),
	}); err != nil || !smsResp.OK {
		rsp.Body = APIErrorCustom("ERROR:CODE:ERROR_VERIFY_CODE", "验证码错误").String()
	} else {
		// smsResp.OK
		rsp.Body = APISuccess(D{
			"msg": "validate success",
		}).String()
	}

	return nil
}
