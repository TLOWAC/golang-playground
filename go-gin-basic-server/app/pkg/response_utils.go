package pkg

import (
	"module/app/constant"
	"module/app/domain/dto"
)

// response 를 만들때 명시적으로 Null 을 표시하는 용도
func Null() interface{} {
	return nil
}

// constant.ResponseStatus 를 활용하여 Response 생성
func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

// 직접 statusCode 와 message 를 사용하여 Response 생성
func BuildResponse_[T any](status string, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
