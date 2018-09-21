package handler

import (
	"iunite.club/api/navo/dto"
)

func APISuccess(data interface{}) *dto.APIResponse {
	return &dto.APIResponse{
		Code:    1,
		SubCode: "default.void.success",
		Message: "接口请求成功",
		Data:    data,
	}
}

func APIError(msg string) *dto.APIResponse {
	return &dto.APIResponse{
		Code:    -100,
		SubCode: "err",
		Message: msg,
	}
}
