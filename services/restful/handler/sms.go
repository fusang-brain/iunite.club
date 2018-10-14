package handler

import (
	"context"

	"github.com/micro/go-micro/client"

	go_api "github.com/emicklei/go-restful"

	sms "iunite.club/services/message/proto/sms"

	"strconv"
)

type SMSHandler struct {
	BaseHandler
	smsService sms.SMSService
}

func NewSMSHandler(c client.Client) *SMSHandler {
	return &SMSHandler{
		smsService: sms.NewSMSService(SMSService, c),
	}
}

func (s *SMSHandler) SendVerifyCode(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	smsService := s.smsService

	params := struct {
		Mobile string `json:"mobile,omitempty" validate:"nonzero"`
		Type   string `json:"type,omitempty" validate:"nonzero"`
	}{}

	if err := s.Bind(req, &params); err != nil {
		ErrorResponse(rsp, s.Error().InternalServerError(err.Error()))
		return
	}

	if err := s.Validate(&params); err != nil {
		ErrorResponse(rsp, s.Error().BadRequest(err.Error()))
		return
	}

	if smsResp, err := smsService.SendVerifyCode(ctx, &sms.SendVerifyCodeRequest{
		Mobile: params.Mobile,
		Type:   params.Type,
	}); err != nil {
		// rsp.Body = APIError(err.Error()).String()
		ErrorResponse(rsp, err)
		return
	} else {
		// rsp.Body = smsResp.
		SuccessResponse(rsp, D{
			"msgBody": smsResp.Message,
		})
	}

	return
}

func (s *SMSHandler) ValidateSimpleCode(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	smsService := s.smsService

	params := struct {
		Mobile string `json:"mobile,omitempty" validate:"nonzero"`
		Code   int    `json:"code,omitempty" validate:"nonzero"`
	}{}

	if err := s.Bind(req, &params); err != nil {
		ErrorResponse(rsp, s.Error().InternalServerError(err.Error()))
		return
	}

	if err := s.Validate(&params); err != nil {
		ErrorResponse(rsp, s.Error().BadRequest(err.Error()))
		return
	}

	if smsResp, err := smsService.ValidateMobileCode(ctx, &sms.ValidateMobileCodeRequest{
		Mobile: params.Mobile,
		Code:   strconv.Itoa(params.Code),
	}); err != nil || !smsResp.OK {
		ErrorResponse(rsp, APIErrorCustom("ERROR:CODE:ERROR_VERIFY_CODE", "验证码错误"))
	} else {
		// smsResp.OK
		SuccessResponse(rsp, D{
			"msg": "validate success",
		})
	}

	return
}
