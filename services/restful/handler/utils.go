package handler

import (
	restful "github.com/emicklei/go-restful"
	"github.com/micro/go-micro/errors"
	"iunite.club/services/restful/dto"
)

type D map[string]interface{}

func SuccessResponse(resp *restful.Response, data interface{}) {
	// resp.Body = APISuccess(data).String()
	resp.WriteAsJson(APISuccess(data))

	return
}

func SuccessResponseWithPage(rsp *restful.Response, page, limit, total int64, list interface{}) {
	// rsp.Body = APISuccess(D{
	// 	"CurrentPage": page,
	// 	"PageSize":    limit,
	// 	"PageTotal":   total,
	// 	"Total":       total,
	// 	"List":        list,
	// }).String()

	rsp.WriteAsJson(APISuccess(D{
		"CurrentPage": page,
		"PageSize":    limit,
		"PageTotal":   total,
		"Total":       total,
		"List":        list,
	}))

	return
}

func ErrorResponse(resp *restful.Response, data ...interface{}) {
	dataLen := len(data)

	if dataLen <= 0 {
		// resp.Body = APIErrorCustom("InteralServerError", "服务器出错了!").String()
		// resp.StatusCode = 500

		resp.WriteAsJson(APIErrorCustom("InteralServerError", "服务器出错了!"))
		resp.WriteHeader(500)
		return
	}

	if dataLen == 1 {
		if err, ok := data[0].(error); ok {
			// resp.Body = ParseAndReturnAPIError(err).String()
			resp.WriteAsJson(ParseAndReturnAPIError(err))
		} else if errString, ok := data[0].(string); ok {
			// resp.Body = APIError(errString).String()
			resp.WriteAsJson(APIError(errString))
		}

		return
	}

	gerr := data[0]
	gmsg := data[1]

	if err, ok := gerr.(error); ok {
		s := "ERR"
		if msg, mok := gmsg.(string); mok {
			s = msg

		}
		// resp.Body = ParseAndReturnAPIError(err, s).String()
		resp.WriteAsJson(ParseAndReturnAPIError(err, s))
		return
	}

	resp.WriteError(500, errors.InternalServerError("iunite.club.srv.navo", "InternalServerError"))
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
