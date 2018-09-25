package handler

import (
	go_api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
	"iunite.club/services/navo/dto"
)

type D map[string]interface{}

func SuccessResponse(resp *go_api.Response, data interface{}) error {
	resp.Body = APISuccess(data).String()

	return nil
}

func ErrorResponse(resp *go_api.Response, data ...interface{}) error {
	dataLen := len(data)

	if dataLen <= 0 {
		resp.Body = APIErrorCustom("InteralServerError", "服务器出错了!").String()
		resp.StatusCode = 500

		return nil
	}

	if dataLen == 1 {
		if err, ok := data[0].(error); ok {
			resp.Body = ParseAndReturnAPIError(err).String()
		} else if errString, ok := data[0].(string); ok {
			resp.Body = APIError(errString).String()
		}

		return nil
	}

	gerr := data[0]
	gmsg := data[1]

	if err, ok := gerr.(error); ok {
		s := "ERR"
		if msg, mok := gmsg.(string); mok {
			s = msg

		}
		resp.Body = ParseAndReturnAPIError(err, s).String()

		return nil
	}

	return errors.InternalServerError("iunite.club.srv.navo", "InternalServerError")
}

func APISuccess(data interface{}) *dto.APIResponse {
	return &dto.APIResponse{
		Code:    1,
		SubCode: "default.void.success",
		Message: "接口请求成功",
		Data:    data,
	}
}

func ParseAndReturnAPIError(err error, msg ...string) *dto.APIResponse {
	e := errors.Parse(err.Error())
	if e.Code == 0 {
		return &dto.APIResponse{
			Code:    -9000,
			SubCode: "Error:Unknown",
			Message: "未知错误",
		}
	}

	resp := &dto.APIResponse{
		Code:    -9001,
		SubCode: e.Detail,
	}

	if len(msg) > 0 {
		resp.Message = msg[0]
	}

	return resp
}

func APIError(msg string) *dto.APIResponse {
	return &dto.APIResponse{
		Code:    -100,
		SubCode: "err",
		Message: msg,
	}
}

func APIErrorCustom(subCode, message string, data ...interface{}) *dto.APIResponse {
	r := &dto.APIResponse{
		Code:    -100,
		SubCode: subCode,
		Message: message,
		// Data:    data[0],
	}

	if len(data) > 0 {
		r.Data = data[0]
	}

	return r
}
